package entity

import (
	model "app/graph/model"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Todos []*Todo
}

func ToModelUser(u *User) *model.User {
	var todos []*model.Todo
	for i := 0; i < len(u.Todos); i++ {
		todos = append(todos, &model.Todo{
			ID:   string(rune(u.Todos[i].ID)),
			Text: u.Todos[i].Text,
			Done: u.Todos[i].Done,
		})
	}

	return &model.User{
		ID:    string(rune(u.ID)),
		Name:  u.Name,
		Todos: todos,
	}
}

func ToModelUsers(us []*User) []*model.User {
	var users []*model.User
	for i := 0; i < len(us); i++ {
		users = append(users, ToModelUser(us[i]))
	}

	return users
}
