package models

type Engineer struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Regions []string `json:"regions"`
}
