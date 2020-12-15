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
	query := `INSERT INTO Links (Title, Address, UserID) VALUES ($1, $2, $3) RETURNING ID`

	err := conn.QueryRow(
		ctx,
		query,
		l.Title, l.Address, l.User.ID,
	).Scan(&l.ID)

	if err != nil {
		return 0, err
	}

	return l.ID, nil
}

func GetAll(ctx context.Context, conn *pgx.Conn) ([]*Link, error) {
	query := `
SELECT
	L.ID, L.Title, L.Address, U.ID, U.Username
FROM Links L
	INNER JOIN Users U ON L.UserID = U.ID`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return []*Link{}, err
	}

	links := []*Link{}
	for rows.Next() {
		link := &Link{}
		link.User = &users.User{}
		err := rows.Scan(
			&link.ID,
			&link.Title,
			&link.Address,
			&link.User.ID,
			&link.User.Username,
		)
		if err != nil {
			return []*Link{}, err
		}
		links = append(links, link)
	}

	return links, nil
}
