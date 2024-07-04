-- name: InsertBranch :one
INSERT INTO Branches (Name, Location, OpeningTime, ClosingTime)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateBranches :exec
UPDATE Branches
SET Name = $2, Location = $3, OpeningTime = $4, ClosingTime = $5
WHERE ID = $1;

-- name: DeleteBranches :exec
DELETE FROM Branches
WHERE ID = $1;

-- name: GetBranches :many
SELECT * FROM Branches;

-- name: GetBranch :one
SELECT * FROM Branches
WHERE ID = $1;