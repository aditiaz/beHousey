package authdto

type RegisterRequest struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ListAs   string `json:"listAs"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address" gorm:"type: text"`
	Image    string `json:"image"`
}

type UpdateImageRequest struct {
	Image string `json:"image"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
