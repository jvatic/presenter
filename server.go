package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	logger := log15.New()

	api := &httpAPI{
		logger:        logger,
		isDevelopment: os.Getenv("ENVIRONMENT") == "development",
	}

	r := httprouter.New()
	r2 := httprouter.New()
	r.GET("/api/slides", api.GetSlides)
	r.NotFound = r2.ServeHTTP
	r2.GET("/*path", api.ServeTemplateOrAsset)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3232"
	}

	logger.Info(fmt.Sprintf("Listening on port %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatal(err)
	}
}

type assetManifest struct {
	Assets map[string]string `json:"assets"`
}

type httpAPI struct {
	logger        log15.Logger
	isDevelopment bool
	assetManifest *assetManifest
}

func (api *httpAPI) Asset(path string) (io.ReadSeeker, error) {
	data, err := Asset(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (api *httpAPI) AssetManifest() (*assetManifest, error) {
	if api.assetManifest != nil {
		return api.assetManifest, nil
	}

	data, err := api.Asset(filepath.Join("build", "manifest.json"))
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(data)
	var manifest *assetManifest
	if err := dec.Decode(&manifest); err != nil {
		return nil, err
	}

	if !api.isDevelopment {
		api.assetManifest = manifest
	}

	return manifest, nil
}

type ByNumber []Slide

func (a ByNumber) Len() int           { return len(a) }
func (a ByNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNumber) Less(i, j int) bool { return a[i].Number < a[j].Number }

type Slide struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	Body   string `json:"body"`
}
type Slides struct {
	Slides []Slide `json:"slides"`
}

func (api *httpAPI) GetSlides(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	manifest, err := api.AssetManifest()
	if err != nil {
		api.logger.Debug(err.Error())
		w.WriteHeader(500)
		return
	}

	res := &Slides{
		Slides: make([]Slide, 0),
	}

	for name, path := range manifest.Assets {
		if strings.HasPrefix(name, "slides/") {
			var data bytes.Buffer
			r, err := api.Asset(filepath.Join("build", path))
			if err != nil {
				api.logger.Debug(err.Error(), "asset", name)
				w.WriteHeader(500)
				return
			}
			if _, err := io.Copy(&data, r); err != nil {
				api.logger.Debug(err.Error(), "asset", name)
				w.WriteHeader(500)
				return
			}
			body := data.String()
			var title string
			s := bufio.NewScanner(&data)
			for s.Scan() {
				line := s.Text()
				if strings.HasPrefix(line, "# ") {
					title = strings.TrimPrefix(line, "# ")
					break
				}
			}
			if err := s.Err(); err != nil {
				api.logger.Debug(err.Error(), "asset", name)
				w.WriteHeader(500)
				return
			}
			if err != nil {
				api.logger.Debug(err.Error(), "asset", name)
				w.WriteHeader(500)
				return
			}
			number, err := strconv.Atoi(strings.TrimSuffix(filepath.Base(name), filepath.Ext(name)))
			res.Slides = append(res.Slides, Slide{
				Title:  title,
				Number: number,
				Body:   body,
			})
		}
	}

	sort.Sort(ByNumber(res.Slides))

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		api.logger.Debug(err.Error())
		w.WriteHeader(500)
	}
}

func (api *httpAPI) ServeTemplateOrAsset(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	path := params.ByName("path")
	if strings.HasPrefix(path, "/assets/") {
		api.ServeAsset(w, req, params)
		return
	}
	api.ServeTemplate(w, req, params)
}

func (api *httpAPI) ServeAsset(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	buildPath := strings.TrimPrefix(params.ByName("path"), "/assets")
	path := filepath.Join("build", buildPath)
	data, err := api.Asset(path)
	if err != nil {
		api.logger.Debug(err.Error())
		w.WriteHeader(404)
		return
	}
	http.ServeContent(w, req, path, time.Now(), data)
}

func (api *httpAPI) ServeTemplate(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	manifest, err := api.AssetManifest()
	if err != nil {
		api.logger.Debug(err.Error())
		return
	}

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Cache-Control", "max-age=0")

	err = htmlTemplate.Execute(w, &htmlTemplateData{
		Assets:        manifest.Assets,
		IsDevelopment: api.isDevelopment,
		Title:         "Presenter",
	})
	if err != nil {
		api.logger.Debug(err.Error())
		return
	}
}
