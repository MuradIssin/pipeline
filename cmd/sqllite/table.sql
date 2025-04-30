-- Таблица для хранения информации о филиалах
CREATE TABLE branches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- Таблица для хранения информации о исполнителях
CREATE TABLE executors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT,
    mobile TEXT
);

-- Таблица для хранения информации о кредитных программах
CREATE TABLE credit_programs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- Таблица для хранения информации о статусах заявок
CREATE TABLE statuses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- Таблица для хранения целей кредита
CREATE TABLE loan_purposes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

-- Основная таблица для проектов
CREATE TABLE projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company TEXT NOT NULL,
    branch_id INTEGER,
    executor_id INTEGER,
    amount INTEGER,
    status_id INTEGER,
    comments TEXT,
    FOREIGN KEY (branch_id) REFERENCES branches(id),
    FOREIGN KEY (executor_id) REFERENCES executors(id),
    FOREIGN KEY (status_id) REFERENCES statuses(id)
);

-- Таблица для связи одного проекта с множеством целей кредита
CREATE TABLE project_loan_purposes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- уникальный идентификатор связи
    project_id INTEGER NOT NULL,           -- ссылка на проект
    purpose_id INTEGER NOT NULL,           -- ссылка на цель кредита
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (purpose_id) REFERENCES loan_purposes(id)
);

-- Таблица для связи одного проекта с множеством кредитных программ
CREATE TABLE project_credit_programs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,      -- уникальный идентификатор связи
    project_id INTEGER NOT NULL,               -- ссылка на проект
    credit_program_id INTEGER NOT NULL,        -- ссылка на кредитную программу
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (credit_program_id) REFERENCES credit_programs(id)
);
