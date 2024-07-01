package models

import (
	db "gitlab/go-prolog-api/example/db/sqlc"
	"time"
)

type CreateProdukRequest struct {
	NamaProduk string `json:"nama_produk" binding:"required"`
	Harga      string `json:"harga" binding:"required"`
	Stok       int32  `json:"stok" binding:"required"`
	Kategori   string `json:"kategori" binding:"required"`
}

type ProdukResponse struct {
	NamaProduk string    `json:"nama_produk"`
	Harga      string    `json:"harga"`
	Stok       int32     `json:"stok"`
	Kategori   string    `json:"kategori"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewProdukResponse(produk db.Produk) ProdukResponse {
	return ProdukResponse{
		NamaProduk: produk.NamaProduk,
		Harga:      produk.Harga,
		Stok:       produk.Stok,
		Kategori:   produk.Kategori,
		CreatedAt:  produk.CreatedAt,
	}
}

type UpdateProdukRequest struct {
	Harga    string `json:"harga" binding:"required"`
	IDProduk int32  `json:"id_produk" binding:"required"`
}
