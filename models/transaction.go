package models

type Transaction struct {
	ID         int      `json:"id"`
	CheckIn    string   `json:"check_in"`
	CheckOut   string   `json:"check_out"`
	PropertyID int      `json:"property_id"`
	Property   Property `json:"property"`
	UserID     int      `json:"user_id"`
	User       User     `json:"user"`
	Price      int      `json:"price"`
	Status     string   `json:"status"`
}

func (Transaction) TableName() string {
	return "transactions"
}
