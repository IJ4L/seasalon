-- name: InsertReservation :one
INSERT INTO reservations (IDUser, IDService)
VALUES ($1, $2) RETURNING *;

-- name: SelectReservation :many
SELECT r.ID, s.Name AS ServiceName, s.Pricing, r.CreatedAt
FROM reservations r
JOIN services s ON r.IDService = s.ID
WHERE r.IDUser = $1;
