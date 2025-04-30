package models

import "time"

type Links struct {
	ID         uint      `json:"id"`
	Link       string    `json:"link"`
	Short_URL  string    `json:"short_url"`
	Created_at time.Time `json:"created_at"`
}
