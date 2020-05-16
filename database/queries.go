package database

import(
    "github.com/jinzhu/gorm"
)

func GetTodoListsAndItems(db *gorm.DB) ([]*TodoList, []*Item) {
	var todoLists []*TodoList
	var todoListsItems []*Item
	db.Find(&todoLists)
	var todoListIds []int
	for _, todoList := range todoLists {
		todoListIds = append(todoListIds, todoList.ID)
	}
	db.Where("todo_list_id IN (?)", todoListIds).Find(&todoListsItems)

    return todoLists, todoListsItems
}

func GetTodoListAndItems(db *gorm.DB, id int) (*TodoList, []*Item) {
	var todoList TodoList
	var todoListItems []*Item
	db.Where("id = ?", id).First(&todoList)
	db.Where("todo_list_id = ?", id).Find(&todoListItems)

    return &todoList, todoListItems
}

func CreateTodoList(db *gorm.DB, name string) *TodoList {
    todoList := TodoList{Name: name}

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
