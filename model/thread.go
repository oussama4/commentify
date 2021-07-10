package model

type Thread struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Domain string `json:"domain"`
	Title  string `json:"title"`
}
