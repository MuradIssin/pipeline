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
		{ID: 1, Name: "Актау"},
		{ID: 2, Name: "Актобе"},
		{ID: 3, Name: "Алматы"},
		{ID: 4, Name: "Астана"},
		{ID: 5, Name: "Атырау"},
		{ID: 6, Name: "Караганда"},
		{ID: 7, Name: "Кокшетау"},
		{ID: 8, Name: "Костанай"},
		{ID: 9, Name: "Кызылорда"},
		{ID: 10, Name: "Павлодар"},
		{ID: 11, Name: "Петропавловск"},
		{ID: 12, Name: "Семей"},
		{ID: 13, Name: "Тараз"},
		{ID: 14, Name: "Уральск"},
		{ID: 15, Name: "Усть-Каменогорск"},
		{ID: 16, Name: "Шымкент"},
	}
	LoanPurposes = []LoanPurpose{
		{ID: 1, Name: "ПОС"},
		{ID: 2, Name: "Инвестиции"},
		{ID: 3, Name: "Гарантии"},
		{ID: 4, Name: "Овердрафт"},
		{ID: 5, Name: "Рефинансирование"},
	}
	CreditPrograms = []CreditProgram{
		{ID: 1, Name: "СС"},
		{ID: 2, Name: "ДО"},
		{ID: 3, Name: "ЕНПФ"},
		{ID: 4, Name: "КЖК"},
		{ID: 5, Name: "Кен Дала"},
		{ID: 6, Name: "Кен Дала 2"},
		{ID: 7, Name: "Астана Бизнес 2"},
	}
	Statuses = []Status{
		{ID: 1, Name: "Сбор документов"},
		{ID: 2, Name: "Индикатив"},
		{ID: 3, Name: "в работе КП"},
		{ID: 4, Name: "в ДФК"},
		{ID: 5, Name: "в ДКР"},
		{ID: 6, Name: "Одобрен"},
		{ID: 7, Name: "Выдан"},
		{ID: 8, Name: "Отказ Клиента"},
		{ID: 9, Name: "Отказ Банка"},
	}
	Executors = []Executor{
		{ID: 1, Name: "Абдыраимов Г.Н."},
		{ID: 2, Name: "Акбердиева З."},
		{ID: 3, Name: "Амиргалиев А/"},
		{ID: 4, Name: "Ахметов Н.Д."},
		{ID: 5, Name: "Баймбетова А.Ж."},
		{ID: 6, Name: "Даулетбаев Д.К."},
		{ID: 7, Name: "Дудка С."},
		{ID: 8, Name: "Ерғали Ж."},
		{ID: 9, Name: "Жамелова А."},
		{ID: 10, Name: "Жанәбілов Б.А"},
		{ID: 11, Name: "Жолдикараева А."},
		{ID: 12, Name: "Зекен А.Е."},
		{ID: 13, Name: "Ибрагимов Г.C."},
		{ID: 14, Name: "Ильясов К."},
		{ID: 15, Name: "Искаков А.А."},
		{ID: 16, Name: "Кабиден А."},
		{ID: 17, Name: "Кожанулы Е."},
		{ID: 18, Name: "Қабдығали А"},
		{ID: 19, Name: "Куанов А."},
		{ID: 20, Name: "Кудайберген А."},
		{ID: 21, Name: "Макажанов О.М."},
		{ID: 22, Name: "Медетова А."},
		{ID: 23, Name: "Муслимов Д.Ш."},
		{ID: 24, Name: "Муканов А."},
		{ID: 25, Name: "Муканов А.Е."},
		{ID: 26, Name: "Нагимова А."},
		{ID: 27, Name: "Найманов Р.С."},
		{ID: 28, Name: "Налекова Л.Б."},
		{ID: 29, Name: "Новикова"},
		{ID: 30, Name: "Нүсүпқұл М."},
		{ID: 31, Name: "Омарбекова А.Р."},
		{ID: 32, Name: "Оразова Б.С."},
		{ID: 33, Name: "Оспанов К"},
		{ID: 34, Name: "Өтел Е.Н."},
		{ID: 35, Name: "Рагулин А"},
		{ID: 36, Name: "Рахимберли О."},
		{ID: 37, Name: "Серик Б.М."},
		{ID: 38, Name: "Супекова Д."},
		{ID: 39, Name: "Ташев Б."},
		{ID: 40, Name: "Тимон А."},
		{ID: 41, Name: "Ткаченко М."},
		{ID: 42, Name: "Тлебалдинова М."},
		{ID: 43, Name: "Уахитова Л.А."},
		{ID: 44, Name: "Устьянцева И.В"},
		{ID: 45, Name: "Вязникова Е.А."},
		{ID: 46, Name: "Шайхы Б."},
		{ID: 47, Name: "Шегирбаев М.С."},
	}
}
