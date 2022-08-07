package entity

import (
	model "app/graph/model"
	"time"
)

type User struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func ToModelUser(u *User) *model.User {
	return &model.User{
		ID:   string(rune(u.ID)),
		Name: u.Name,
	}
}

func ToModelUsers(us []*User) []*model.User {
	var users []*model.User
	for i := 0; i < len(us); i++ {
		users = append(users, ToModelUser(us[i]))
	}

	return users
}
