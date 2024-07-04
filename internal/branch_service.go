package server

import (
	db "gitlab/go-prolog-api/example/db/sqlc"
	"gitlab/go-prolog-api/example/internal/models"
	"gitlab/go-prolog-api/example/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateBranch(ctx *gin.Context) {
	var req models.CreateBranchRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.InsertBranchParams{
		Name:        req.Name,
		Location:    req.Location,
		Openingtime: req.Openingtime,
		Closingtime: req.Closingtime,
	}

	branch, err := server.repo.InsertBranch(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	rsp := models.SuccessResponse{
		Status:  "success",
		Message: "success create branch",
		Data:    models.NewBranchResponse(branch),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetBranchs(ctx *gin.Context) {
	branchs, err := server.repo.GetBranches(ctx)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	rsp := models.SuccessResponse{
		Status:  "success",
		Message: "success get services",
		Data:    models.NewBranchsResponse(branchs),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetBranch(ctx *gin.Context) {
	param := ctx.Query("branch_id")
	produkID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	branch, err := server.repo.GetBranch(ctx, int32(produkID))
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	rsp := models.SuccessResponse{
		Status:  "success",
		Message: "success get branch by id",
		Data:    models.NewBranchResponse(branch),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) DestroyBranch(ctx *gin.Context) {
	param := ctx.Query("branch_id")
	branchID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	_, err = server.repo.GetBranch(ctx, int32(branchID))
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	err = server.repo.DeleteBranches(ctx, int32(branchID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus branch"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "branch berhasil dihapus"})
}

func (server *Server) UpdateBranch(ctx *gin.Context) {
	var req models.UpdateBranchsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	_, err := server.repo.GetBranch(ctx, int32(req.ID))
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	arg := db.UpdateBranchesParams{
		ID:          int32(req.ID),
		Name:        req.Name,
		Location:    req.Location,
		Openingtime: req.Openingtime,
		Closingtime: req.Closingtime,
	}

	err = server.repo.UpdateBranches(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "branch berhasil diupdate"})
}
