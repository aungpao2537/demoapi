package models

import "time"

type GetUserResponse struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
type MessageResponse struct {
	Message string `json:"message"`
}
