package model

import "github.com/oussama4/commentify/base/validate"

type Page struct {
	Id    string `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type PageInput struct {
	Url, Title string
}

func (pi *PageInput) Valid() error {
	v := validate.New()

	v.Check(pi.Title != "", "title", "page title is required")
	v.Check(pi.Url != "", "url", "page url is required")

	return v.Valid()
}
