package models

import "database/sql"

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
}

type ProjectModel struct {
	DB *sql.DB
}

// import "database/sql"

// // Структура для проекта

// // This will insert a new snippet into the database.
// func (m *ProjectModel) Insert(p Project) (int, error) {

// 	// Начинаем транзакцию
// 	tx, err := m.Begin()
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer tx.Rollback() // Откатим транзакцию, если что-то пойдет не так

// 	// Вставляем проект
// 	res, err := tx.Exec(`INSERT INTO projects (company, branch_id, executor_id, amount, status_id, comments)
//                          VALUES (?, ?, ?, ?, ?, ?)`,
// 		p.Company, p.BranchID, p.ExecutorID, p.Amount, p.StatusID, p.Comments)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Получаем ID вставленного проекта
// 	projectID, err := res.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Вставляем связи с целями кредитования
// 	for _, purposeID := range p.LoanPurposeIDs {
// 		_, err := tx.Exec("INSERT INTO project_loan_purposes (project_id, purpose_id) VALUES (?, ?)", projectID, purposeID)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	// Вставляем связи с кредитными программами
// 	for _, programID := range p.CreditProgramIDs {
// 		_, err := tx.Exec("INSERT INTO project_credit_programs (project_id, credit_program_id) VALUES (?, ?)", projectID, programID)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	// Если все прошло успешно, коммитим транзакцию
// 	if err := tx.Commit(); err != nil {
// 		return 0, err
// 	}

// 	// Возвращаем ID вставленного проекта
// 	return projectID, nil
// }

// // This will return a specific snippet based on its id.
// func (m *ProjectModel) Get(id int) (*Project, error) {
// 	return nil, nil
// }

// // This will return the 10 most recently created snippets.
// func (m *ProjectModel) Latest() ([]*Project, error) {
// 	return nil, nil
// }

func (m *ProjectModel) createTables() ([]*Project, error) {
	return nil, nil
}
