package persistence

import (
	"database/sql"
	"log"
	"os"
	"task-api/internal/domain"
	"task-api/internal/port/outbound"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type mysqlRepo struct {
	db *sql.DB
}

// ConnectToMysql : To establish a connection to our mysql database
func ConnectToMysql() *sql.DB {
	errs := godotenv.Load("cmd/.env")
	if errs != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = os.Getenv("DBNET")
	cfg.Addr = os.Getenv("DBADR")
	cfg.DBName = "taskdb"
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	return db
}

// CallMysql : To set the value of db and make it a outbound interface
func CallMysql(db *sql.DB) outbound.Database {
	return mysqlRepo{db: db}
}

func (r mysqlRepo) CreateTask(task domain.Task) (string, error) {
	id := uuid.New().String()
	_, err := r.db.Exec("INSERT INTO tasks (id,title,description) VALUES (?,?,?)", id, task.Title, task.Description)
	return id, err
}

func (r mysqlRepo) GetByID(id string) (domain.Task, error) {
	row := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = ?", id)

	var task domain.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r mysqlRepo) GetAll() ([]domain.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r mysqlRepo) UpdateTask(task domain.Task) error {
	row := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = ?", task.ID)

	var tasks domain.Task
	err := row.Scan(&tasks.ID, &tasks.Title, &tasks.Description)
	if err != nil {
		return err
	}
	_, errs := r.db.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", task.Title, task.Description, task.ID)
	return errs
}

func (r mysqlRepo) Delete(id string) error {
	row := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = ?", id)

	var tasks domain.Task
	err := row.Scan(&tasks.ID, &tasks.Title, &tasks.Description)
	if err != nil {
		return err
	}

	query := "DELETE FROM tasks WHERE id = ?"
	_, errs := r.db.Exec(query, id)
	return errs
}
