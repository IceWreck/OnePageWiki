package handlers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/IceWreck/OnePageWiki/config"
	"github.com/IceWreck/OnePageWiki/logger"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// ViewTemplateData is the data we pass to the view html template
type ViewTemplateData struct {
	WikiTitle    string
	MarkdownHTML template.HTML
}

// MarkdownView - view rendered markdown
func MarkdownView(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open(config.MarkdownLocation)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			logger.Error(err)
		}
	}()

	fileMarkdown, err := ioutil.ReadAll(file)
	if err != nil {
		fileMarkdown = []byte("### Error Reading File")
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert(fileMarkdown, &buf); err != nil {
		logger.Error(err)
	}

	err = templateCache["view"].Execute(w, ViewTemplateData{
		WikiTitle:    config.WikiTitle,
		MarkdownHTML: template.HTML(buf.String()),
	})
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
