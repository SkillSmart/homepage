package domain

// Define the sub interface to domain.User
type User struct {
	Username string `json:"Username"`
	AuthToken string `json:"auth_token`
	UserId uint64 `json:"user_id"`
}

// Public interface
func (u *User) New() *User {
	//TODO: Intialize new User on creation
	return &User{}
}

// Handling function
func (u *User) IsValid() bool {
	//TODO: Implement check to validate User account within system
	return true
}
