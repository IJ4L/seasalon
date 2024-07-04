package server

import (
	db "gitlab/go-prolog-api/example/db/sqlc"
	"gitlab/go-prolog-api/example/internal/models"
	"gitlab/go-prolog-api/example/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateReservation(ctx *gin.Context) {
	var req models.CreateReservationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.InsertReservationParams{
		Iduser:    req.Iduser,
		Idservice: req.Idservice,
	}

	reservation, err := server.repo.InsertReservation(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	user, err := server.repo.GetUserByID(ctx, reservation.Iduser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	service, err := server.repo.GetService(ctx, reservation.Idservice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	rsp := models.SuccessResponse{
		Status:  "success",
		Message: "success create reservation",
		Data:    models.NewReservationResponse(reservation, user, service),
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) GetReservations(ctx *gin.Context) {
	param := ctx.Query("user_id")
	userID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	reservations, err := server.repo.SelectReservation(ctx, int32(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	user, err := server.repo.GetUserByID(ctx, int32(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	service, err := server.repo.GetService(ctx, reservations[0].ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	rsp := models.SuccessResponse{
		Status:  "success",
		Message: "success get services",
		Data:    models.NewReservationsResponse(reservations, user, service),
	}

	ctx.JSON(http.StatusOK, rsp)
}
