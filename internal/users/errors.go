package users

type WrongUsernameOrPasswordError struct{}

func (w *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
