package data

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

type Executor struct {
	ID     int
	Name   string
	Email  string
	Mobile string
}

// Глобальные срезы для справочных данных
var Branches []Branch
var LoanPurposes []LoanPurpose
var CreditPrograms []CreditProgram
var Statuses []Status
var Executors []Executor

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
	Executors = []Executor{
		{ID: 1, Name: "Murad"},
		{ID: 2, Name: "Ivan"},
		{ID: 3, Name: "Kuanish"},
	}
}
