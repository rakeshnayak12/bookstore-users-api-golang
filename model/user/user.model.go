package user

// User defines attribute of user
type User struct {
	UserID    int    `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
