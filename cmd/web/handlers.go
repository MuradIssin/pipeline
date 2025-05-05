package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pipeline/internal/models"
	"strconv"
	"text/template"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w) // Use the notFound() helper
		return
	}
	projects, err := app.projects.AllIn()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, project := range projects {
		fmt.Fprintf(w, "%+v\n", project)
	}
	// files := []string{
	// 	"./ui/html/base.html",
	// 	"./ui/html/pages/home.html",
	// 	"./ui/html/partials/nav.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal Server Error", 500)
	// 	app.serverError(w, err) // Use the serverError() helper.
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	// http.Error(w, "Internal Server Error", 500)
	// 	app.serverError(w, err) // Use the serverError() helper.
	// }
}

func (app *application) pipeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/partials/view.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", project)
	if err != nil {
		app.serverError(w, err)
	}

	// fmt.Fprintf(w, "%+v", project)
}

func (app *application) pipeCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}

	project := models.Project{
		Company:    "ООО Ромашка",
		BranchID:   1,
		ExecutorID: 2,
		Amount:     5000000,
		StatusID:   1,
		Comments:   "Приоритетный проект",
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

	http.Redirect(w, r, fmt.Sprintf("/pipe/view?id=%d", projectID), http.StatusSeeOther)

}
