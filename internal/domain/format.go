package domain

// Task : To set the format and structure of database
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
