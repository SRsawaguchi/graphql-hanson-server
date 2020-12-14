package links

import (
	"context"

	"github.com/SRsawaguchi/graphql-hanson-server/internal/users"
	"github.com/jackc/pgx/v4"
)

type Link struct {
	ID      int64
	Title   string
	Address string
	User    *users.User
}

func (l *Link) Save(ctx context.Context, conn *pgx.Conn) (int64, error) {
	query := `INSERT INTO Links (Title, Address) VALUES ($1, $2) RETURNING ID`

	err := conn.QueryRow(
		ctx,
		query,
		l.Title, l.Address,
	).Scan(&l.ID)

	if err != nil {
		return 0, err
	}

	return l.ID, nil
}
