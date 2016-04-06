package main

import (
	"log"

	matrix "github.com/jvatic/asset-matrix-go"
)

func main() {
	m := matrix.New(&matrix.Config{
		Paths: []*matrix.AssetRoot{
			{
				Path: "./src",
			},
			{
				Path: "./vendor",
			},
			{
				GitRepo:   "git://github.com/jvatic/marbles-js.git",
				GitBranch: "master",
				GitRef:    "6057b0600550667e37e0e320e8bddc6563292139",
				Path:      "src",
			},
			{
				GitRepo:   "git://github.com/chjj/marked.git",
				GitBranch: "master",
				GitRef:    "88ce4df47c4d994dc1b1df1477a21fb893e11ddc",
				Path:      "lib",
			},
		},
		Outputs: []string{
			"presenter.scss",
			"presenter.js",
			"react.prod.js",
			"react-dom.prod.js",
			"react.dev.js",
			"react-dom.dev.js",
			"marked.js",
			"highlight.js",
			"highlight/github.css",
			"slides/*.md",
		},
		OutputDir: "./build",
	})
	if err := m.Build(); err != nil {
		log.Fatal(err)
	}
	m.RemoveOldAssets()
}
