-- name: CreateUser :one
INSERT INTO
    users (
        fullname, photo, email, phonenumber, password, role
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: UpdatePassword :exec
UPDATE users SET password = $2 WHERE id = $1;

-- name: UpdatePositionAndUsername :exec
UPDATE users SET fullname = $2, role = $3 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
