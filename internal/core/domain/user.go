package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}

func NewPerson(email, password, name string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) GetEmail() string {
	return u.Email
}
