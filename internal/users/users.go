package users

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
