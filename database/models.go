package database

import(
    "github.com/jinzhu/gorm"
)

type Config struct {
    DatabaseType string
    Hostname string
    Username string
    Password string
    Database string

}

type TodoList struct {
    ID       int    `json:"id" gorm:"primary_key"`
    Name     string `json:"name"`
    Items []*Item `gorm:"foreignkey:TodoListId"`
}

type Item struct {
    ID       int    `json:"id" gorm:"primary_key"`
    TodoListId int `json:"todo_list_id"`
    Name     string `json:"name"`
    DueDate  string    `json:"due_date"`
}

func CreateTodoList(db *gorm.DB, name string, items ...*Item) *TodoList {
    todoList := TodoList{Name: name, Items: items}

    db.NewRecord(todoList)

    db.Create(&todoList)

    return &todoList

}

func CreateItem(db *gorm.DB, todoListId int, name string, dueDate string) *Item {
    item := Item{TodoListId: todoListId, Name: name, DueDate: dueDate}

    db.NewRecord(item)

    db.Create(&item)

    return &item

}
