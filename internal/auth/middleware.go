package auth

import (
	"context"
	"net/http"

	"github.com/SRsawaguchi/graphql-hanson-server/internal/users"
	"github.com/SRsawaguchi/graphql-hanson-server/pkg/jwt"
	"github.com/jackc/pgx/v4"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware(conn *pgx.Conn) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(r.Context(), conn, username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
