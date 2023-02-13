package authdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	ListAs   string `json:"listAs"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address" gorm:"type: text"`
	Image    string `json:"image"`
}