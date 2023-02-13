package authdto

type RegisterResponse struct {
	Username string ` json:"username"`
	Message  string ` json:"message"`
	ListAs   string `gorm:"type: int" json:"listAs"`
}

type LoginResponse struct {
	Username string ` json:"username"`
	Token    string ` json:"token"`
	ListAs   string `json:"listAs"`
}
