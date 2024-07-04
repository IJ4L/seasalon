-- name: InsertRating :one
INSERT INTO ratings (IDUser, IDService, Rating, Comment)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetRatings :many
SELECT r.Rating, r.Comment, r.CreatedAt, u.Fullname AS UserName
FROM ratings r
JOIN users u ON r.IDUser = u.ID
WHERE r.IDService = $1;
