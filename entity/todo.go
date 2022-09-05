package entity

import (
	model "app/graph/model"
	"strconv"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID uint64 `json:"userId"`
	User   *User
}

func ToModelTodo(t *Todo) *model.Todo {
	return &model.Todo{
		ID:   strconv.Itoa(int(t.ID)),
		Text: t.Text,
		Done: t.Done,
		User: ToModelUser(t.User),
	}
}

func ToModelTodos(ts []*Todo) []*model.Todo {
	var todos []*model.Todo
	for i := 0; i < len(ts); i++ {
		todos = append(todos, ToModelTodo(ts[i]))
	}

	return todos
}
