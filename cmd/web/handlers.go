package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pipeline/internal/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w) // Use the notFound() helper
	// 	return
	// }

	projects, err := app.projects.AllIn()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Projects = projects
	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) pipeView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}
	project, err := app.projects.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Project = project
	app.render(w, http.StatusOK, "view.html", data)

	// app.render(w, http.StatusOK, "view.html", &templateData{
	// 	Project: project,
	// })

	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/partials/nav.html",
	// 	"./ui/html/partials/view.html",
	// }

	// // Регистрируем функцию и парсим шаблоны
	// funcMap := template.FuncMap{
	// 	"GetBranchName":    models.GetBranchName,
	// 	"GetUserName":      models.GetUserName,
	// 	"GetGoalsName":     models.GetCreditGoal,
	// 	"FormatNumberView": models.FormatNumber,
	// 	"GetCreditName":    models.GetCreditProg,
	// 	"GetStatusName":    models.GetStatus,
	// 	"FDate":            models.FormatDate,
	// }
	// tmpl, err := template.New("base").Funcs(funcMap).ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// data := &templateData{
	// 	Project: project,
	// }
	// err = tmpl.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) pipeCreate(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display the form for creating a new snippet..."))
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "create.html", data)
}

func (app *application) pipeCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	company := r.PostForm.Get("company")
	comments := r.PostForm.Get("comments")

	// expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	// if err != nil {
	// 		app.clientError(w, http.StatusBadRequest)
	// 		return
	// }

	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// 	app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
	// 	return
	// }

	project := models.Project{
		Company:    company,
		BranchID:   1,
		ExecutorID: 2,
		Amount:     5000001,
		StatusID:   1,
		Comments:   comments,
		LoanPurposeIDs: []int{
			1, // Пополнение оборотных средств
			2, // Закупка сырья
		},
		CreditProgramIDs: []int{
			1, // Кредит на развитие
			2, // Кредит для малого бизнеса
		},
		LastUpdate: time.Now(),
	}

	// Вставляем проект в базу
	projectID, err := app.projects.Insert(project)
	if err != nil {
		log.Fatal(err)
	}
	// Выводим ID вставленного проекта
	fmt.Println("Проект успешно добавлен с ID:", projectID)

	// w.Write([]byte("Create a new project..."))
	// Redirect the user to the relevant page for the snippet.

	http.Redirect(w, r, fmt.Sprintf("/pipe/view/%d", projectID), http.StatusSeeOther)

}
