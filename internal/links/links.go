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

func GetAll(ctx context.Context, conn *pgx.Conn) ([]*Link, error) {
	query := `SELECT ID, Title, Address FROM Links`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return []*Link{}, err
	}

	links := []*Link{}
	for rows.Next() {
		link := &Link{}
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			return []*Link{}, err
		}
		links = append(links, link)
	}

	return links, nil
}
