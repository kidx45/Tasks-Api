package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"task-api/internal/domain"
	"task-api/internal/port/outbound"

	"github.com/google/uuid"

	//imported to use psql
	"github.com/joho/godotenv"
)

type postgresRepo struct {
	db *sql.DB
}

// ConnectToPostgres : To estblish a connection with our postgres database
func ConnectToPostgres() *sql.DB {
	errs := godotenv.Load("cmd/.env")
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	location := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		"task_api",
	)

	db, err := sql.Open("postgres", location)
	if err != nil {
		log.Fatal(err)
	}

	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	return db
}

// CallPsql : To set the value of db and make it an outbound interface
func CallPsql(db *sql.DB) outbound.Database {
	return postgresRepo{db: db}
}

func (r postgresRepo) CreateTask(c context.Context, task domain.Task) (string, error) {
	id := uuid.New().String()
	_, err := r.db.ExecContext(c, "INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3)", id, task.Title, task.Description)
	return id, err
}

func (r postgresRepo) GetByID(c context.Context, id string) (domain.Task, error) {
	row := r.db.QueryRowContext(c, "SELECT id, title, description FROM tasks WHERE id = $1", id)

	var task domain.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r postgresRepo) GetAll(c context.Context) ([]domain.Task, error) {
	rows, err := r.db.QueryContext(c, "SELECT id, title, description FROM tasks")
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

func (r postgresRepo) UpdateTask(c context.Context, task domain.Task) error {
	row := r.db.QueryRowContext(c, "SELECT id, title, description FROM tasks WHERE id = $1", task.ID)

	var tasks domain.Task
	err := row.Scan(&tasks.ID, &tasks.Title, &tasks.Description)
	if err != nil {
		return err
	}
	_, errs := r.db.ExecContext(c, "UPDATE tasks SET title = $1, description = $2 WHERE id = $3", task.Title, task.Description, task.ID)
	return errs
}

func (r postgresRepo) Delete(c context.Context, id string) error {
	row := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = $1", id)

	var tasks domain.Task
	err := row.Scan(&tasks.ID, &tasks.Title, &tasks.Description)
	if err != nil {
		return err
	}

	query := "DELETE FROM tasks WHERE id = $1"
	_, errs := r.db.Exec(query, id)
	return errs
}
