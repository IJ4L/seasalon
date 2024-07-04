package server

import (
	db "gitlab/go-prolog-api/example/db/sqlc"
	"gitlab/go-prolog-api/example/internal/models"
	"gitlab/go-prolog-api/example/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateService(ctx *gin.Context) {
	var req models.CreateServiceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.InsertServicesParams{
		Idbranch: req.Idbranch,
		Name:     req.Name,
		Type:     req.Type,
		Detail:   req.Detail,
		Pricing:  req.Pricing,
	}

	service, err := server.repo.InsertServices(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.NewServiceResponse(service))
}

func (server *Server) GetAllService(ctx *gin.Context) {
	services, err := server.repo.GetServices(ctx)
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
		Data:    models.NewServicesResponse(services),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetService(ctx *gin.Context) {
	param := ctx.Query("service_id")
	produkID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	service, err := server.repo.GetService(ctx, int32(produkID))
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
		Message: "success get service",
		Data:    models.NewServiceResponse(service),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) DestroyService(ctx *gin.Context) {
	param := ctx.Query("service_id")
	produkID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	_, err = server.repo.GetService(ctx, int32(produkID))
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	err = server.repo.DeleteServices(ctx, int32(produkID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus service"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "service berhasil dihapus"})
}

func (server *Server) UpdateService(ctx *gin.Context) {
	var req models.UpdateServicesRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	_, err := server.repo.GetService(ctx, int32(req.ID))
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	arg := db.UpdateServicesParams{
		ID:       req.ID,
		Name:     req.Name,
		Type:     req.Type,
		Detail:   req.Detail,
		Pricing:  req.Pricing,
		Duration: req.Duration,
	}

	err = server.repo.UpdateServices(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "service berhasil diupdate"})
}
