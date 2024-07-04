// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (
        fullname, photo, email, phonenumber, password, role
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, fullname, photo, email, phonenumber, password, role, createdat, updatedat
`

type CreateUserParams struct {
	Fullname    string `json:"fullname"`
	Photo       string `json:"photo"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Fullname,
		arg.Photo,
		arg.Email,
		arg.Phonenumber,
		arg.Password,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Photo,
		&i.Email,
		&i.Phonenumber,
		&i.Password,
		&i.Role,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, fullname, photo, email, phonenumber, password, role, createdat, updatedat FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Photo,
		&i.Email,
		&i.Phonenumber,
		&i.Password,
		&i.Role,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, fullname, photo, email, phonenumber, password, role, createdat, updatedat FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Photo,
		&i.Email,
		&i.Phonenumber,
		&i.Password,
		&i.Role,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users SET password = $2 WHERE id = $1
`

type UpdatePasswordParams struct {
	ID       int32  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.ID, arg.Password)
	return err
}

const updatePositionAndUsername = `-- name: UpdatePositionAndUsername :exec
UPDATE users SET fullname = $2, role = $3 WHERE id = $1
`

type UpdatePositionAndUsernameParams struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
}

func (q *Queries) UpdatePositionAndUsername(ctx context.Context, arg UpdatePositionAndUsernameParams) error {
	_, err := q.db.ExecContext(ctx, updatePositionAndUsername, arg.ID, arg.Fullname, arg.Role)
	return err
}
