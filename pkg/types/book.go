package types

import "time"

type Book struct {
	Id              int32      `json:"id"`
	Title           string     `json:"title"`
	Author          string     `json:"author"`
	PublicationDate *time.Time `json:"publication_date"`
	Quantity        int32      `json:"quantity"`
	Genre           string     `json:"genre"`
	Description     string     `json:"description"`
	Rating          float32    `json:"rating"`
	Address         string     `json:"address"`
}
