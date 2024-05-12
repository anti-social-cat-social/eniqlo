package customer

import "time"

type Customer struct {
	ID        	string     	`json:"id" db:"id"`
	Name      	string    	`json:"name"`
	PhoneNumber	string		`json:"phoneNumber" db:"phone_number"`
	CreatedAt	time.Time 	`json:"createdAt" db:"created_at"`
}

type CustomerRegisterDTO struct {
	Name     	string `json:"name" validate:"required,min=5,max=50"`
	PhoneNumber	string `json:"phoneNumber" validate:"required,min=10,max=16"`
}

type CustomerRegisterResponse struct {
	UserId		string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"` 
	Name 		string `json:"name"`
}