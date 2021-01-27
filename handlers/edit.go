package handlers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/IceWreck/OnePageWiki/config"
	"github.com/IceWreck/OnePageWiki/logger"
)

// EditTemplateData is the data we pass to the edit html template
type EditTemplateData struct {
	WikiTitle string
	Markdown  string
}

// EditView - view rendered markdown
func EditView(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open(config.MarkdownLocation)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			logger.Error(err)
		}
	}()

	md, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	err = templateCache["edit"].Execute(w, EditTemplateData{
		WikiTitle: config.WikiTitle,
		Markdown:  string(md),
	})
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// EditForm - save input to file
func EditForm(w http.ResponseWriter, r *http.Request) {
	updatedMarkdown := r.PostFormValue("markdown-input")

	file, err := os.Create(config.MarkdownLocation)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			logger.Error(err)
		}
	}()

	_, err = file.WriteString(updatedMarkdown)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("markdown file updated")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
