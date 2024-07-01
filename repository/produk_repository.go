package repository

import (
	"context"
	"database/sql"
	"fmt"
	db "gitlab/go-prolog-api/example/db/sqlc"
)

type Produk interface {
	db.Querier
}

type SQLProduk struct {
	*db.Queries
	db *sql.DB
}

func NewProduk(dbms *sql.DB) Produk {
	return &SQLProduk{
		db:      dbms,
		Queries: db.New(dbms),
	}
}

func (SQLProduk *SQLProduk) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := SQLProduk.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}