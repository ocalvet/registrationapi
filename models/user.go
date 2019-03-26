package models

// User Who is registering?
type User struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Skills    []string `json:"skills"`
}
