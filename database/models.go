package database

type Config struct {
	DatabaseType string
	Hostname     string
	Username     string
	Password     string
	Database     string
}

type TodoList struct {
	ID    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Items []*Item `gorm:"foreignkey:TodoListId"`
}

type Item struct {
	ID         int    `json:"id" gorm:"primary_key"`
	TodoListId int    `json:"todo_list_id"`
	Name       string `json:"name"`
	DueDate    string `json:"due_date"`
}
