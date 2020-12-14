package users

import (
	"context"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Create(ctx context.Context, conn *pgx.Conn) (int, error) {
	query := `INSERT INTO Users (Username, Password) VALUES($1, $2) RETURNING ID, Password`

	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return 0, err
	}
	err = conn.QueryRow(
		ctx,
		query,
		u.Username, hashedPassword,
	).Scan(&u.ID, &u.Password)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(ctx context.Context, conn *pgx.Conn, username string) (int, error) {
	query := `SELECT ID FROM Users WHERE Username = $1`

	var id int
	err := conn.QueryRow(
		ctx,
		query,
		username,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
