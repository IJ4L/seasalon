-- name: InsertServices :one
INSERT INTO services (IDBranch, Name, Type, Detail, Pricing, Duration)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateServices :exec
UPDATE services
SET Name = $2, Type = $3, Detail = $4, Pricing = $5, Duration = $6, UpdatedAt = CURRENT_TIMESTAMP
WHERE ID = $1;

-- name: DeleteServices :exec
DELETE FROM services
WHERE ID = $1;

-- name: GetServices :many
SELECT * FROM services;

-- name: GetService :one
SELECT * FROM services
WHERE ID = $1;

-- name: GetServicesByBranch :many
SELECT * FROM services
WHERE Type = $1;