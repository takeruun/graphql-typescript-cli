package entity

import (
	model "app/graph/model"
	"time"
)

type Todo struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	Text      string `json:"text"`
	Done      bool   `json:"done"`
	UserID    uint64 `json:"userId"`
	User      *User
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func ToModelTodo(t *Todo) *model.Todo {
	return &model.Todo{
		ID:   string(rune(t.ID)),
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
