package pg

import (
	"context"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New(dbURL string) (*pgxpool.Pool, sqlc.Querier) {
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	return pool, sqlc.New(pool)
}
