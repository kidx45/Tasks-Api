package domain

// Task : To set the format and structure of database
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UserInput : TO accept user input when creating task
type UserInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
