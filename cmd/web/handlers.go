package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pipeline/internal/data"
	"pipeline/internal/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	projects, err := app.projects.AllIn()
	if err != nil {
		app.serverError(w, err)
		return
	}
	dataForPage := app.newTemplateData(r)
	dataForPage.Projects = projects
	app.render(w, http.StatusOK, "home.html", dataForPage)
}

func (app *application) pipeView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
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

	dataForPage := app.newTemplateData(r)
	dataForPage.Project = project
	app.render(w, http.StatusOK, "view.html", dataForPage)
}

func (app *application) pipeCreate(w http.ResponseWriter, r *http.Request) {
	dataForPage := app.newTemplateData(r)
	dataForPage.Branches = data.Branches
	dataForPage.Executors = data.Executors
	dataForPage.LoanPurposes = data.LoanPurposes
	dataForPage.CreditPrograms = data.CreditPrograms
	dataForPage.Statuses = data.Statuses
	app.render(w, http.StatusOK, "create.html", dataForPage)
}

func (app *application) pipeCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	company := r.PostForm.Get("company")

	branchStr := r.PostForm.Get("branch")
	if branchStr == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	branchID, err := strconv.Atoi(branchStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	executorStr := r.PostForm.Get("executor")
	if executorStr == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	executorId, err := strconv.Atoi(executorStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	loanPurposeIDsStr := r.Form["LoanPurposes"] // []string
	var loanPurposeIDs []int
	for _, idStr := range loanPurposeIDsStr {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			loanPurposeIDs = append(loanPurposeIDs, id)
		}
	}

	creditProgramIDsStr := r.Form["CreditPrograms"] // []string
	var creditProgramIDs []int
	for _, idStr := range creditProgramIDsStr {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			creditProgramIDs = append(creditProgramIDs, id)
		}
	}

	amountStr := r.PostForm.Get("amount")
	if amountStr == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	amountInt, err := strconv.Atoi(amountStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	statusStr := r.PostForm.Get("status")
	if statusStr == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	statusID, err := strconv.Atoi(statusStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	comments := r.PostForm.Get("comments")

	project := models.Project{
		Company:          company,
		BranchID:         branchID,
		ExecutorID:       executorId,
		LoanPurposeIDs:   loanPurposeIDs,
		CreditProgramIDs: creditProgramIDs,
		Amount:           uint(amountInt),
		StatusID:         statusID,
		Comments:         comments,
		LastUpdate:       time.Now(),
	}

	// Вставляем проект в базу
	projectID, err := app.projects.Insert(project)
	if err != nil {
		log.Fatal(err)
	}
	// Выводим ID вставленного проекта
	fmt.Println("Проект успешно добавлен с ID:", projectID)
	app.infoLog.Println("dd")
	// app.infoLog.Sprintf("Проект успешно добавлен с ID:", projectID)

	http.Redirect(w, r, fmt.Sprintf("/pipe/view/%d", projectID), http.StatusSeeOther)
}
