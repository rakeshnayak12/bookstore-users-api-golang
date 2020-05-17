package user

// User defines attribute of user
type User struct {
	UserID    uint64 `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
