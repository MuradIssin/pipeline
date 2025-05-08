package main

import (
	"path/filepath"
	"pipeline/internal/data"
	"pipeline/internal/models"
	"text/template"
)

type templateData struct {
	CurrentYear    int
	Project        *models.Project
	Projects       []*models.Project
	Branches       []data.Branch // ✅ добавили это поле
	Executors      []data.Executor
	LoanPurposes   []data.LoanPurpose
	CreditPrograms []data.CreditProgram
	Statuses       []data.Status
	Form           any
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
		// update
		ts, err := template.New(name).Funcs(funcMap).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		// Call ParseFiles() *on this template set* to add the  page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
