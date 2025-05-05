package main

// import (
// 	"database/sql"
// 	"fmt"
// )

// // Структура для филиала
// type Branch struct {
// 	ID   int
// 	Name string
// }

// // Структура для исполнителя
// type Executor struct {
// 	ID     int
// 	Name   string
// 	Email  string
// 	Mobile string
// }

// // Структура для кредитной программы
// type CreditProgram struct {
// 	ID   int
// 	Name string
// }

// // Структура для статуса заявки
// type Status struct {
// 	ID   int
// 	Name string
// }

// // Структура для цели кредита
// type LoanPurpose struct {
// 	ID   int
// 	Name string
// }

// type Project struct {
// 	ID               int
// 	Company          string
// 	BranchID         int
// 	ExecutorID       int
// 	Amount           uint
// 	StatusID         int
// 	Comments         string
// 	LoanPurposeIDs   []int
// 	CreditProgramIDs []int
// }

// type ProjectModel struct {
// 	DB *sql.DB
// }

// // func insertProject(db *sql.DB, p Project) (int64, error) {
// // 	// Начинаем транзакцию
// // 	tx, err := db.Begin()
// // 	if err != nil {
// // 		return 0, err
// // 	}
// // 	defer tx.Rollback() // Откатим транзакцию, если что-то пойдет не так

// // 	// Вставляем проект
// // 	res, err := tx.Exec(`INSERT INTO projects (company, branch_id, executor_id, amount, status_id, comments)
// //                          VALUES (?, ?, ?, ?, ?, ?)`,
// // 		p.Company, p.BranchID, p.ExecutorID, p.Amount, p.StatusID, p.Comments)
// // 	if err != nil {
// // 		return 0, err
// // 	}

// // 	// Получаем ID вставленного проекта
// // 	projectID, err := res.LastInsertId()
// // 	if err != nil {
// // 		return 0, err
// // 	}

// // 	// Вставляем связи с целями кредитования
// // 	for _, purposeID := range p.LoanPurposeIDs {
// // 		_, err := tx.Exec("INSERT INTO project_loan_purposes (project_id, purpose_id) VALUES (?, ?)", projectID, purposeID)
// // 		if err != nil {
// // 			return 0, err
// // 		}
// // 	}

// // 	// Вставляем связи с кредитными программами
// // 	for _, programID := range p.CreditProgramIDs {
// // 		_, err := tx.Exec("INSERT INTO project_credit_programs (project_id, credit_program_id) VALUES (?, ?)", projectID, programID)
// // 		if err != nil {
// // 			return 0, err
// // 		}
// // 	}

// // 	// Если все прошло успешно, коммитим транзакцию
// // 	if err := tx.Commit(); err != nil {
// // 		return 0, err
// // 	}

// // 	// Возвращаем ID вставленного проекта
// // 	return projectID, nil
// // }

// func createTables(db *sql.DB) error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS branches (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL
// 	);

// 	CREATE TABLE IF NOT EXISTS executors (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		email TEXT,
// 		mobile TEXT
// 	);

// 	CREATE TABLE IF NOT EXISTS credit_programs (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL
// 	);

// 	CREATE TABLE IF NOT EXISTS statuses (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL
// 	);

// 	CREATE TABLE IF NOT EXISTS loan_purposes (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL
// 	);

// 	CREATE TABLE IF NOT EXISTS projects (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		company TEXT NOT NULL,
// 		branch_id INTEGER,
// 		executor_id INTEGER,
// 		amount INTEGER,
// 		status_id INTEGER,
// 		comments TEXT,
// 		FOREIGN KEY (branch_id) REFERENCES branches(id),
// 		FOREIGN KEY (executor_id) REFERENCES executors(id),
// 		FOREIGN KEY (status_id) REFERENCES statuses(id)
// 	);

// 	-- Таблица для связи одного проекта с множеством целей кредита
// 	CREATE TABLE IF NOT EXISTS project_loan_purposes (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,  -- уникальный идентификатор связи
// 		project_id INTEGER NOT NULL,           -- ссылка на проект
// 		purpose_id INTEGER NOT NULL,           -- ссылка на цель кредита
// 		FOREIGN KEY (project_id) REFERENCES projects(id),
// 		FOREIGN KEY (purpose_id) REFERENCES loan_purposes(id)
// 	);

// 	-- Таблица для связи одного проекта с множеством кредитных программ
// 	CREATE TABLE IF NOT EXISTS project_credit_programs (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,      -- уникальный идентификатор связи
// 		project_id INTEGER NOT NULL,               -- ссылка на проект
// 		credit_program_id INTEGER NOT NULL,        -- ссылка на кредитную программу
// 		FOREIGN KEY (project_id) REFERENCES projects(id),
// 		FOREIGN KEY (credit_program_id) REFERENCES credit_programs(id)
// 	);

// 	-- Индексы для улучшения производительности
// 	CREATE INDEX IF NOT EXISTS idx_project_date ON projects(id);
// 	CREATE INDEX IF NOT EXISTS idx_project_loan_purposes_project_id ON project_loan_purposes(project_id);
// 	CREATE INDEX IF NOT EXISTS idx_project_credit_programs_project_id ON project_credit_programs(project_id);
// 	`

// 	// Выполнение запроса на создание таблиц
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("ошибка при создании таблиц: %v", err)
// 	}

// 	return nil
// }
