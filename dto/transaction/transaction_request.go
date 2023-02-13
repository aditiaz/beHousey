package transactiondto

import (
	"housey/models"
)

type CreateTransactionRequest struct {
	CheckIn    string    `json:"check_in"`
	CheckOut   string    `json:"check_out"`
	PropertyID    int          `json:"property_id" form:"property_id" `
	Property      models.Property `json:"property"`
	UserID     int          `json:"user_id" form:"user_id" `
	User       models.User  `json:"user"`
	Price      int     `json:"price"  `
	Status     string       `json:"status" form:"status" `

}

type UpdateTransactionRequest struct {
	Status string `json:"status" form:"status"`
}
