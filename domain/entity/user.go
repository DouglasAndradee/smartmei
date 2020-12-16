package entity

// UserFetch -
type UserFetch interface {
}

// User - It's a model of user in the system
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
