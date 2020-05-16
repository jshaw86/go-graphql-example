package transformations

import (
    "strconv"
	"github.com/jshaw86/go-graphql-example/graph/model"
	"github.com/jshaw86/go-graphql-example/database"
)

func ToGraphQLTodoList(todoList *database.TodoList, items []*database.Item) (*model.TodoList, error) {
    todoListItems, _ := ToGraphQLItems(items)
    todoListModel := model.TodoList{
        Name: &todoList.Name,
        Items: todoListItems,

    }

    return &todoListModel, nil

}

func ToGraphQLTodoLists(todoLists []*database.TodoList, items []*database.Item) ([]*model.TodoList, error) {
	var graphQLTodoLists []*model.TodoList
    itemsByTodoListId := make(map[int][]*database.Item)

    for _, item := range items {
        if _, ok := itemsByTodoListId[item.TodoListId]; !ok {
            itemsByTodoListId[item.TodoListId] = make([]*database.Item, 0)
        }

        itemsByTodoListId[item.TodoListId] = append(itemsByTodoListId[item.TodoListId], item)

    }

	for _, todoList := range todoLists {
        todoListItems :=  itemsByTodoListId[todoList.ID]
		graphQLTodoList, _ := ToGraphQLTodoList(todoList, todoListItems)
		graphQLTodoLists = append(graphQLTodoLists, graphQLTodoList)

	}
	return graphQLTodoLists, nil
}

func ToGraphQLItems(items []*database.Item) ([]*model.Item, error) {
    var graphQLItems []*model.Item

    for _, item := range items {
        graphQLItem, _ := ToGraphQLItem(item)
        graphQLItems = append(graphQLItems, graphQLItem)

    }

    return graphQLItems, nil
}

func ToGraphQLItem(item *database.Item) (*model.Item, error) {
    itemModel := model.Item{
        ID: strconv.Itoa(item.ID),
        TodoListID: strconv.Itoa(item.TodoListId),
        Name: item.Name,
        DueDate: item.DueDate,
    }

    return &itemModel, nil
}
