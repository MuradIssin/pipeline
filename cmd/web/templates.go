package main

import (
	"path/filepath"
	"pipeline/internal/models"
	"text/template"
)

type templateData struct {
	Project  *models.Project
	Projects []*models.Project
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	funcMap := template.FuncMap{
		"GetBranchName":    models.GetBranchName,
		"GetUserName":      models.GetUserName,
		"GetGoalsName":     models.GetCreditGoal,
		"FormatNumberView": models.FormatNumber,
		"GetCreditName":    models.GetCreditProg,
		"GetStatusName":    models.GetStatus,
		"FDate":            models.FormatDate,
	}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			page,
		}
		// ts, err := template.ParseFiles(files...)
		ts, err := template.New(name).Funcs(funcMap).ParseFiles(files...)

		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
