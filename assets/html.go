package assets

import (
	"embed"
	"html/template"
	"os"
)

//go:embed app.js screen.css
var Assets embed.FS

//go:embed index.gohtml
var templates embed.FS

var funcMap = template.FuncMap{
	"hostname": os.Hostname,
}

var Index = template.Must(template.New("index.gohtml").Funcs(funcMap).ParseFS(templates, "index.gohtml"))
