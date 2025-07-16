package models

type Task struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"` //due-date eg:2025-03-04
	Status string `json:"status"`

}