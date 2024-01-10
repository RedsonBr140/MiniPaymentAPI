package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	DB *pgx.Conn
	Queries *Queries
}

func Initialize(ctx *context.Context, pgUrl string) (*Postgres, error) {
	conn, err := pgx.Connect(*ctx, pgUrl)
	if err != nil {
		return &Postgres{}, nil
	}

	queries := New(conn)

	return &Postgres{
		DB: conn,
		Queries: queries,
	}, nil
}