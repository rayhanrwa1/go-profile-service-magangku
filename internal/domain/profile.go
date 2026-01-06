package domain

type Profile struct {
	UserID      string `json:"user_id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
