-- name: CreateProduk :one
INSERT INTO
    Produk (nama_produk, harga, stok, kategori)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetAllProduk :many
SELECT * FROM Produk;

-- name: GetProduk :one
SELECT FROM Produk WHERE id_produk = $1;

-- name: UpdateProduk :exec
UPDATE Produk SET harga = $1 WHERE id_produk = $2 RETURNING *;

-- name: DeleteProduk :exec
DELETE FROM Produk WHERE id_produk = $1;