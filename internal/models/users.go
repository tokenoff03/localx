package models

type SuperUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// TODO: check table fields
type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"pfp"`
	Description    string `json:"description"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	City           string `json:"city"`
	Instagram      string `json:"instagram"`
	Interests      string `json:"interests"`
}
