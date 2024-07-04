package repository

import (
	"context"
	"database/sql"
	"fmt"
	db "gitlab/go-prolog-api/example/db/sqlc"
)

type Repo interface {
	db.Querier
}

type SQLRepo struct {
	*db.Queries
	db *sql.DB
}

func NewRepos(dbms *sql.DB) Repo {
	return &SQLRepo{
		db:      dbms,
		Queries: db.New(dbms),
	}
}

func (SQLProduk *SQLRepo) execTx(ctx context.Context, fn func(*db.Queries) error) error {
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