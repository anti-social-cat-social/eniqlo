package customer

import (
	"time"
)

type Customer struct {
	ID          string
	PhoneNumber string `db:"phone_number"`
	Name        string
	CreatedAt   time.Time `db:"created_at"`
}

type QueryParams struct {
	Limit       int    `form:"limit"`
	Offset      int    `form:"offset"`
	PhoneNumber string `form:"phoneNumber"`
	Name        string `form:"name"`
}
