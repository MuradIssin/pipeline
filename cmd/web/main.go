// go build -o pipeline ./cmd/web

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pipeline/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	projects *models.ProjectModel
}

const fileDb = "pipeline.db"

type Branch struct {
	ID   int
	Name string
}

type LoanPurpose struct {
	ID   int
	Name string
}

type CreditProgram struct {
	ID   int
	Name string
}

type Status struct {
	ID   int
	Name string
}

// Глобальные срезы для справочных данных
var Branches []Branch
var LoanPurposes []LoanPurpose
var CreditPrograms []CreditProgram
var Statuses []Status

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// запрашиваем подключение к БД - проверяем наличия файла

	appPath, err := os.Executable()
	if err != nil {
		errorLog.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), fileDb)
	// Проверяем, существует ли файл базы данных
	_, err = os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
		infoLog.Println("need make new db", fileDb, install)
	}
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		infoLog.Fatal("error when db open")
		return
	}
	defer db.Close()
	if install {
		err = models.CreateTables(db)
		// err = models.CreateTables(db)
		if err != nil {
			infoLog.Fatal("Ошибка при создании таблиц:", err)
		}
		infoLog.Println("база данных создана", fileDb)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		projects: &models.ProjectModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func init() {
	Branches = []Branch{
		{ID: 1, Name: "Алматы"},
		{ID: 2, Name: "Астана"},
		{ID: 3, Name: "Шымкент"},
	}

	LoanPurposes = []LoanPurpose{
		{ID: 1, Name: "Пополнение оборотных средств"},
		{ID: 2, Name: "Приобретение оборудования"},
		{ID: 3, Name: "Расширение бизнеса"},
	}

	CreditPrograms = []CreditProgram{
		{ID: 1, Name: "Программа 1"},
		{ID: 2, Name: "Программа 2"},
		{ID: 3, Name: "Программа 3"},
	}

	Statuses = []Status{
		{ID: 1, Name: "В процессе"},
		{ID: 2, Name: "Одобрено"},
		{ID: 3, Name: "Отклонено"},
	}
}
