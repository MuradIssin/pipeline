package models

import (
	"database/sql"
	"fmt"
	"pipeline/internal/data"
	"strconv"
	"time"
)

type Project struct {
	ID               int
	Company          string
	BranchID         int
	ExecutorID       int
	Amount           uint
	StatusID         int
	Comments         string
	LoanPurposeIDs   []int
	CreditProgramIDs []int
	Created          time.Time
	LastUpdate       time.Time
}

type ProjectModel struct {
	DB *sql.DB
}

func CreateTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS branches (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS executors (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT,
			mobile TEXT,
			is_deleted BOOLEAN NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS loan_purposes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS credit_programs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS statuses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			company TEXT NOT NULL,
			branch_id INTEGER,
			executor_id INTEGER,
			amount INTEGER,
			status_id INTEGER,
			comments TEXT,
			is_deleted BOOLEAN NOT NULL DEFAULT 0,
			created DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
			last_update DATETIME NOT NULL,
			FOREIGN KEY (branch_id) REFERENCES branches(id),
			FOREIGN KEY (executor_id) REFERENCES executors(id),
			FOREIGN KEY (status_id) REFERENCES statuses(id)
		);

		CREATE TABLE IF NOT EXISTS project_loan_purposes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER NOT NULL,
			purpose_id INTEGER NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT 0,
			FOREIGN KEY (project_id) REFERENCES projects(id),
			FOREIGN KEY (purpose_id) REFERENCES loan_purposes(id)
		);

		CREATE TABLE IF NOT EXISTS project_credit_programs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER NOT NULL,
			credit_program_id INTEGER NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT 0,
			FOREIGN KEY (project_id) REFERENCES projects(id),
			FOREIGN KEY (credit_program_id) REFERENCES credit_programs(id)
		);
`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблиц: %v", err)
	}
	// Вставка начальных значений в таблицы
	err = insertInitialData(db)
	if err != nil {
		return fmt.Errorf("ошибка при вставке начальных данных: %v", err)
	}

	return nil
}

// Функция для вставки начальных данных в справочники
func insertInitialData(db *sql.DB) error {
	// Вставка значений для филиалов
	branches := []string{"Алматы", "Астана", "ШЫмкент"}
	for _, branch := range branches {
		_, err := db.Exec(`INSERT INTO branches (name) VALUES (?)`, branch)
		if err != nil {
			return fmt.Errorf("ошибка при вставке филиала: %v", err)
		}
	}

	// Вставка значений для кредитных программ
	creditPrograms := []string{"Программа 1", "Программа 2", "Программа 3"}
	for _, program := range creditPrograms {
		_, err := db.Exec(`INSERT INTO credit_programs (name) VALUES (?)`, program)
		if err != nil {
			return fmt.Errorf("ошибка при вставке кредитной программы: %v", err)
		}
	}

	// Вставка значений для статусов заявки
	statuses := []string{"В процессе", "Одобрено", "Отклонено"}
	for _, status := range statuses {
		_, err := db.Exec(`INSERT INTO statuses (name) VALUES (?)`, status)
		if err != nil {
			return fmt.Errorf("ошибка при вставке статуса заявки: %v", err)
		}
	}

	// Вставка значений для целей кредитования
	loanPurposes := []string{"Пополнение оборотных средств", "Приобретение оборудования", "Расширение бизнеса"}
	for _, purpose := range loanPurposes {
		_, err := db.Exec(`INSERT INTO loan_purposes (name) VALUES (?)`, purpose)
		if err != nil {
			return fmt.Errorf("ошибка при вставке цели кредитования: %v", err)
		}
	}

	return nil
}

// func LoadBranches(db *sql.DB) ([]Branch, error) {
// 	var branches []Branch
// 	rows, err := db.Query("SELECT id, name FROM branches WHERE is_deleted != 1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var branch Branch
// 		if err := rows.Scan(&branch.ID, &branch.Name); err != nil {
// 			return nil, err
// 		}
// 		branches = append(branches, branch)
// 	}

// 	return branches, nil
// }

// func LoadLoanPurposes(db *sql.DB) ([]LoanPurpose, error) {
// 	var loanPurposes []LoanPurpose
// 	rows, err := db.Query("SELECT id, name FROM loan_purposes WHERE is_deleted != 1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var loanPurpose LoanPurpose
// 		if err := rows.Scan(&loanPurpose.ID, &loanPurpose.Name); err != nil {
// 			return nil, err
// 		}
// 		loanPurposes = append(loanPurposes, loanPurpose)
// 	}

// 	return loanPurposes, nil
// }

// func LoadCreditPrograms(db *sql.DB) ([]CreditProgram, error) {
// 	var creditPrograms []CreditProgram
// 	rows, err := db.Query("SELECT id, name FROM credit_programs WHERE is_deleted != 1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var creditProgram CreditProgram
// 		if err := rows.Scan(&creditProgram.ID, &creditProgram.Name); err != nil {
// 			return nil, err
// 		}
// 		creditPrograms = append(creditPrograms, creditProgram)
// 	}

// 	return creditPrograms, nil
// }

// func LoadStatuses(db *sql.DB) ([]Status, error) {
// 	var statuses []Status
// 	rows, err := db.Query("SELECT id, name FROM statuses WHERE is_deleted != 1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var status Status
// 		if err := rows.Scan(&status.ID, &status.Name); err != nil {
// 			return nil, err
// 		}
// 		statuses = append(statuses, status)
// 	}

// 	return statuses, nil
// }

func (m *ProjectModel) Insert(p Project) (int, error) {
	// Вставка основного проекта
	res, err := m.DB.Exec(`
		INSERT INTO projects (company, branch_id, executor_id, amount, status_id, comments, last_update)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		p.Company, p.BranchID, p.ExecutorID, p.Amount, p.StatusID, p.Comments, p.LastUpdate)
	if err != nil {
		return 0, fmt.Errorf("ошибка при вставке проекта: %v", err)
	}

	projectIDLast, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("не удалось получить ID проекта: %v", err)
	}
	projectID := int(projectIDLast)

	// Вставка целей кредита
	for _, purposeID := range p.LoanPurposeIDs {
		_, err := m.DB.Exec(`
			INSERT INTO project_loan_purposes (project_id, purpose_id)
			VALUES (?, ?)`, projectID, purposeID)
		if err != nil {
			return 0, fmt.Errorf("ошибка при вставке цели кредита: %v", err)
		}
	}

	// Вставка кредитных программ
	for _, creditProgramID := range p.CreditProgramIDs {
		_, err := m.DB.Exec(`
			INSERT INTO project_credit_programs (project_id, credit_program_id)
			VALUES (?, ?)`, projectID, creditProgramID)
		if err != nil {
			return 0, fmt.Errorf("ошибка при вставке кредитной программы: %v", err)
		}
	}

	return projectID, nil
}

// This will return a specific snippet based on its id.
func (m *ProjectModel) Get(id int) (*Project, error) {
	// Получаем основную информацию о проекте
	query := `
	SELECT id, company, branch_id, executor_id, amount, status_id, comments, created, last_update
	FROM projects
	WHERE id = ? AND is_deleted != 1
	`
	row := m.DB.QueryRow(query, id)
	project := &Project{}
	err := row.Scan(
		&project.ID,
		&project.Company,
		&project.BranchID,
		&project.ExecutorID,
		&project.Amount,
		&project.StatusID,
		&project.Comments,
		&project.Created,
		&project.LastUpdate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRecord // проект не найден
		}
		return nil, fmt.Errorf("ошибка при получении проекта: %v", err)
	}

	// Загрузка связанных кредитных программ
	creditProgramIDs, err := m.GetCreditProgramIDs(id)
	if err != nil {
		return nil, err
	}
	project.CreditProgramIDs = creditProgramIDs

	// Загрузка связанных целей кредитования
	loanPurposeIDs, err := m.GetLoanPurposeIDs(id)
	if err != nil {
		return nil, err
	}
	project.LoanPurposeIDs = loanPurposeIDs
	return project, nil
}

// This will return the 10 most recently created snippets.
func (m *ProjectModel) AllIn() ([]*Project, error) {
	stmt := `
	SELECT id, company, branch_id, executor_id, amount, status_id, comments
	FROM projects
	WHERE is_deleted != 1
	ORDER BY id
	`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	projects := []*Project{}
	for rows.Next() {
		p := &Project{}
		err = rows.Scan(
			&p.ID,
			&p.Company,
			&p.BranchID,
			&p.ExecutorID,
			&p.Amount,
			&p.StatusID,
			&p.Comments,
		)
		if err != nil {
			return nil, err
		}

		creditProgramIDs, err := m.GetCreditProgramIDs(p.ID)
		if err != nil {
			return nil, err
		}
		p.CreditProgramIDs = creditProgramIDs

		// Загрузка связанных целей кредитования
		loanPurposeIDs, err := m.GetLoanPurposeIDs(p.ID)
		if err != nil {
			return nil, err
		}
		p.LoanPurposeIDs = loanPurposeIDs

		projects = append(projects, p)
	}
	return projects, nil
}

func (m *ProjectModel) GetCreditProgramIDs(projectID int) ([]int, error) {
	query := `
		SELECT credit_program_id
		FROM project_credit_programs
		WHERE project_id = ? AND is_deleted != 1
	`

	rows, err := m.DB.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func (m *ProjectModel) GetLoanPurposeIDs(projectID int) ([]int, error) {
	query := `
		SELECT purpose_id
		FROM project_loan_purposes
		WHERE project_id = ? AND is_deleted != 1
	`

	rows, err := m.DB.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

// Вспомогательная функция для получения имени филиала по ID
func GetBranchName(id int) string {
	for _, b := range data.Branches {
		if b.ID == id {
			return b.Name
		}
	}
	return "Неизвестный филиал"
}

func GetUserName(id int) string {
	for _, b := range data.Executors {
		if b.ID == id {
			return b.Name
		}
	}
	return "Неизвестное имя"
}

func GetCreditGoal(id int) string {
	for _, b := range data.LoanPurposes {
		if b.ID == id {
			return b.Name
		}
	}
	return "Неизвестная цель кредитования"
}

func GetCreditProg(id int) string {
	for _, b := range data.CreditPrograms {
		if b.ID == id {
			return b.Name
		}
	}
	return "Неизвестная программа кредитования"
}

// Разделение числа на разряды с пробелами (например: 1 000 000)
func FormatNumber(n uint) string {
	s := strconv.FormatUint(uint64(n), 10)

	result := ""
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if count != 0 && count%3 == 0 {
			result = " " + result
		}
		result = string(s[i]) + result
		count++
	}
	return result
}

func GetStatus(id int) string {
	for _, b := range data.Statuses {
		if b.ID == id {
			return b.Name
		}
	}
	return "Неизвестная программа кредитования"
}

func FormatDate(t time.Time) string {
	// return t.Format("02.01.2006 15:04")
	location := time.FixedZone("UTC+5", 5*60*60) // 5 часов * 60 минут * 60 секунд
	return t.In(location).Format("02.01.2006 15:04") + " "
}
