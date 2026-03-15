package domain

type (
	Client struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
	}
)
