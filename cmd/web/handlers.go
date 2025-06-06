package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pipeline/internal/data"
	"pipeline/internal/models"
	"pipeline/internal/validator"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type projectCreateForm struct {
	Expires                   int
	Company                   string
	SelectedBranch            string
	SelectedBranchID          int
	SelectedExecutorId        int
	SelectedLoanPurposesIDs   []int
	SelectedCreditProgramsIDs []int
	Amount                    uint
	SelectedStatusesId        int
	Comment                   string
	// FieldErrors               map[string]
	validator.Validator
}

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

	dataForPage.Form = projectCreateForm{}
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

	// обработчик ошибок
	form := projectCreateForm{
		Company:                   r.PostForm.Get("company"),
		SelectedBranch:            r.PostForm.Get("branch"),
		SelectedBranchID:          branchID,
		SelectedExecutorId:        executorId,
		SelectedLoanPurposesIDs:   loanPurposeIDs,
		SelectedCreditProgramsIDs: creditProgramIDs,
		Amount:                    uint(amountInt),
		SelectedStatusesId:        statusID,
		Comment:                   comments,
		// FieldErrors:               map[string]string{},
	}
	fmt.Println("Company:", company)
	form.CheckField(validator.NotBlank(company), "company", "Компания должна иметь название")
	// Initialize a map to hold any validation errors for the form fields.
	// fieldErrors := make(map[string]string)
	// if strings.TrimSpace(company) == "" {
	// 	form.FieldErrors["company"] = "Компания должна иметь название"
	// 	app.infoLog.Println(form.FieldErrors["company"])
	// } else if utf8.RuneCountInString(company) > 100 {
	// 	form.FieldErrors["company"] = "This field cannot be more than 100 characters long"
	// }
	form.CheckField(validator.MinNum(len(loanPurposeIDs), 1), "LoanPurposes", "нужно выбрать цель")
	// if len(loanPurposeIDs) == 0 {
	// 	form.FieldErrors["LoanPurposes"] = "нужно выбрать цель"
	// 	app.infoLog.Println(form.FieldErrors["LoanPurposes"])
	// }

	form.CheckField(validator.MinNum(len(creditProgramIDs), 1), "CreditPrograms", "нужно выбрать программу")

	// if len(creditProgramIDs) == 0 {
	// 	form.FieldErrors["CreditPrograms"] = "нужно выбрать программу"
	// 	app.infoLog.Println(form.FieldErrors["CreditPrograms"])
	// }

	// if len(form.FieldErrors) > 0 {

	if !form.Validator.Valid() {
		dataForPage := app.newTemplateData(r)
		dataForPage.Project = &models.Project{
			// ID:      id,
			Company: company,
			Amount:  uint(amountInt),
		}
		// dataForPage.Project.Company = company
		dataForPage.Form = form
		dataForPage.Branches = data.Branches
		dataForPage.Executors = data.Executors
		dataForPage.LoanPurposes = data.LoanPurposes
		dataForPage.CreditPrograms = data.CreditPrograms
		// dataForPage.Project.Amount = uint(amountInt)
		dataForPage.Statuses = data.Statuses
		app.infoLog.Println("find errors on forms")
		app.render(w, http.StatusUnprocessableEntity, "create.html", dataForPage)
		return
	}

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
	app.infoLog.Println("Проект успешно добавлен с ID:", projectID)

	http.Redirect(w, r, fmt.Sprintf("/pipe/view/%d", projectID), http.StatusSeeOther)
}

func (app *application) pipeUpdate(w http.ResponseWriter, r *http.Request) {

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

	form := projectCreateForm{
		Company:                   project.Company,
		SelectedBranchID:          project.BranchID,
		SelectedExecutorId:        project.ExecutorID,
		SelectedLoanPurposesIDs:   project.LoanPurposeIDs,
		SelectedCreditProgramsIDs: project.CreditProgramIDs,
		Amount:                    project.Amount,
		SelectedStatusesId:        project.StatusID,
		Comment:                   project.Comments,
		// FieldErrors:               map[string]string{},
	}

	dataForPage := app.newTemplateData(r)
	dataForPage.Form = form
	dataForPage.Project = project

	dataForPage.Branches = data.Branches
	dataForPage.Executors = data.Executors
	dataForPage.LoanPurposes = data.LoanPurposes
	dataForPage.CreditPrograms = data.CreditPrograms
	dataForPage.Statuses = data.Statuses

	app.render(w, http.StatusOK, "edit.html", dataForPage)
}

func (app *application) pipeUpdatePost(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}

	err = r.ParseForm()
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

	// обработчик ошибок
	form := projectCreateForm{
		Company:                   r.PostForm.Get("company"),
		SelectedBranch:            r.PostForm.Get("branch"),
		SelectedBranchID:          branchID,
		SelectedExecutorId:        executorId,
		SelectedLoanPurposesIDs:   loanPurposeIDs,
		SelectedCreditProgramsIDs: creditProgramIDs,
		Amount:                    uint(amountInt),
		SelectedStatusesId:        statusID,
		Comment:                   comments,
		// FieldErrors:               map[string]string{},
	}

	form.CheckField(validator.NotBlank(company), "company", "Компания должна иметь название")
	// Initialize a map to hold any validation errors for the form fields.
	// fieldErrors := make(map[string]string)
	// if strings.TrimSpace(company) == "" {
	// 	form.FieldErrors["company"] = "Компания должна иметь название"
	// 	app.infoLog.Println(form.FieldErrors["company"])
	// } else if utf8.RuneCountInString(company) > 100 {
	// 	form.FieldErrors["company"] = "This field cannot be more than 100 characters long"
	// }

	form.CheckField(validator.MinNum(len(loanPurposeIDs), 1), "LoanPurposes", "нужно выбрать цель")
	// if len(loanPurposeIDs) == 0 {
	// 	form.FieldErrors["LoanPurposes"] = "нужно выбрать цель"
	// 	app.infoLog.Println(form.FieldErrors["LoanPurposes"])
	// }

	form.CheckField(validator.MinNum(len(creditProgramIDs), 1), "CreditPrograms", "нужно выбрать программу")
	// if len(creditProgramIDs) == 0 {
	// 	form.FieldErrors["CreditPrograms"] = "нужно выбрать программу"
	// 	app.infoLog.Println(form.FieldErrors["CreditPrograms"])
	// }

	// if len(form.FieldErrors) > 0 {
	if !form.Valid() {
		dataForPage := app.newTemplateData(r)
		dataForPage.Project = &models.Project{
			ID:      id,
			Company: company,
			Amount:  uint(amountInt),
		}
		// dataForPage.Project = &models.Project{}
		// dataForPage.Project.Company = company
		dataForPage.Form = form
		dataForPage.Branches = data.Branches
		dataForPage.Executors = data.Executors
		dataForPage.LoanPurposes = data.LoanPurposes
		dataForPage.CreditPrograms = data.CreditPrograms
		// dataForPage.Project.Amount = uint(amountInt)
		dataForPage.Statuses = data.Statuses
		app.infoLog.Println("find errors on forms")
		app.render(w, http.StatusUnprocessableEntity, "edit.html", dataForPage)
		return
	}

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

	err = app.projects.Update(id, project)
	if err != nil {
		log.Fatal(err)
	}
	app.infoLog.Println("Проект успешно  обновлен с ID:", id)

	http.Redirect(w, r, fmt.Sprintf("/pipe/view/%d", id), http.StatusSeeOther)

	// w.WriteHeader(http.StatusNotImplemented)
	// w.Write([]byte("Функция редактирования проекта ещё не реализована. POST"))
}

func (app *application) pipeDelete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}

	// Вставляем проект в базу
	err = app.projects.SoftDelete(id)
	if err != nil {
		log.Fatal(err)
	}
	app.infoLog.Println("Проект успешно мягко удален ID:", id)

	// http.Redirect(w, r, fmt.Sprintf("/pipe/view/%d", id), http.StatusSeeOther)
	http.Redirect(w, r, "/", http.StatusSeeOther)

	// w.WriteHeader(http.StatusNotImplemented)
	// w.Write([]byte("Функция удаление проекта ещё не реализована."))
}
