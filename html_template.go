package main

import "text/template"

type htmlTemplateData struct {
	Assets        map[string]string
	IsDevelopment bool
	Title         string
}

var htmlTemplate = template.Must(template.New("presenter.html").Funcs(template.FuncMap{
	"assetPath": func(ctx *htmlTemplateData, name string) string {
		return ctx.Assets[name]
	},
}).Parse(`
<!doctype html>
<html>
<head>
  <title>{{.Title}}</title>
  <link rel="stylesheet" type="text/css" href="/assets/{{assetPath . "highlight/github.css"}}" />
  <link rel="stylesheet" type="text/css" href="/assets/{{assetPath . "presenter.css"}}" />
</head>

<body>
  <div id="main"></div>

{{if .IsDevelopment}}
  <script type="application/javascript" src="/assets/{{assetPath . "react.dev.js"}}"></script>
  <script type="application/javascript" src="/assets/{{assetPath . "react-dom.dev.js"}}"></script>
{{else}}
  <script type="application/javascript" src="/assets/{{assetPath . "react.prod.js"}}"></script>
  <script type="application/javascript" src="/assets/{{assetPath . "react-dom.prod.js"}}"></script>
{{end}}
  <script type="application/javascript" src="/assets/{{assetPath . "highlight.js"}}"></script>
  <script type="application/javascript" src="/assets/{{assetPath . "marked.js"}}"></script>
  <script type="application/javascript" src="/assets/{{assetPath . "presenter.js"}}"></script>

  <noscript>
    <h1>You need JavaScript enabled to use this app.</h1>
  </noscript>
</body>
</html>
`))
