package model

import "time"

type Cake struct {
	Id			string		`json:"id"`
	Title		string		`json:"title"`
	Description	string      `json:"description"`
	Rating		float64		`json:"rating"`
	Image		string		`json:"image"`
	Created_at	string		`json:"created_at"`
	Updated_at 	string	    `json:"updated_at"`
}

type InputCake struct {
	Id			string		`json:"id" `
	Title		string		`json:"title" binding:"required"`
	Description	string      `json:"description" binding:"required"`
	Rating		float64		`json:"rating" binding:"required"`
	Image		string		`json:"image" binding:"required"`
	Created_at	time.Time	`json:"created_at" `
	Updated_at 	time.Time	`json:"updated_at" `
}
