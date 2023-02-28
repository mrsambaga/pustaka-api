package book

import "encoding/json"

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required, number"` // field must be filled & int number. Input can be string or int
	Subtitle string      `json:"sub_title"`
}
