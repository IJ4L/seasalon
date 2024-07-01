package server

import (
	db "gitlab/go-prolog-api/example/db/sqlc"
	"gitlab/go-prolog-api/example/internal/models"
	"gitlab/go-prolog-api/example/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateProduk(ctx *gin.Context) {
	var req models.CreateProdukRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.CreateProdukParams{
		NamaProduk: req.NamaProduk,
		Harga:      req.Harga,
		Stok:       req.Stok,
		Kategori:   req.Kategori,
	}

	produk, err := server.produk.CreateProduk(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	rsp := models.NewProdukResponse(produk)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetAllProduk(ctx *gin.Context) {
	rsp, err := server.produk.GetAllProduk(ctx)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) DestroyProduk(ctx *gin.Context) {
	param := ctx.Query("produk_id")
	produkID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	_, err = server.produk.GetProduk(ctx, int32(produkID))
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	err = server.produk.DeleteProduk(ctx, int32(produkID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}

func (server *Server) UpdateProduk(ctx *gin.Context) {
	var req models.UpdateProdukRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.UpdateProdukParams{
		Harga:    req.Harga,
		IDProduk: req.IDProduk,
	}

	err := server.produk.UpdateProduk(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produk berhasil diupdate"})
}