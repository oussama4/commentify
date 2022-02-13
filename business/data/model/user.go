package model

import "github.com/oussama4/commentify/base/validate"

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Name, Email, Password string
}

func (ui *UserInput) Valid() error {
	v := validate.New()

	v.Check(ui.Name != "", "title", "user name is required")
	v.Check(validate.EmailFormat(ui.Email), "email", "invalid email")

	return v.Valid()
}
