package handlers

import (
	"html/template"

	"github.com/IceWreck/OnePageWiki/logger"
)

var templateCache = cacheTemplates()

// we want to parse and cache templates on program start instead of parsing them per request
func cacheTemplates() map[string](*template.Template) {
	logger.Info("parsing templates")
	cache := map[string](*template.Template){
		"view": template.Must(template.ParseFiles([]string{
			"./templates/layout.html",
			"./templates/view.html",
		}...)),
		"edit": template.Must(template.ParseFiles([]string{
			"./templates/layout.html",
			"./templates/edit.html",
		}...)),
	}
	return cache
}
