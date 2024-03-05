package models

type Counter struct {
	Id   uint   `json:"id"`
	Url  string `json:"url"`
	Code string `json:"code"`
	Name string `json:"name"`
}
