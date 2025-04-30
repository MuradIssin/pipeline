// go build -o pipeline ./cmd/web

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

const fileDb = "pipeline.db"

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
		err = createTables(db)
		if err != nil {
			infoLog.Fatal("Ошибка при создании таблиц:", err)
		}
		infoLog.Println("база данных создана", fileDb)
	}

	// Пример подключения к базе данных SQLite
	database, err := sql.Open("sqlite3", "./project.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Пример использования структур
	project := Project{
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
	}

	// Вставляем проект в базу
	projectID, err := insertProject(database, project)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим ID вставленного проекта
	fmt.Println("Проект успешно добавлен с ID:", projectID)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
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
