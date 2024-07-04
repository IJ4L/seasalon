package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	db "gitlab/go-prolog-api/example/db/sqlc"
	"gitlab/go-prolog-api/example/internal/models"
	"gitlab/go-prolog-api/example/util"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Fullname:    req.Fullname,
		Password:    hashedPassword,
		Role:        req.Role,
		Email:       req.Email,
		Photo:       req.Photo,
		Phonenumber: req.Phone,
	}

	user, err := server.repo.CreateUser(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	rsp := models.NewUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	user, err := server.repo.GetUser(ctx, req.Email)
	if err != nil {
		if errors.Is(err, util.ErrRecordNotFound) {
			err := errors.New("akun tidak ditemukan")
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		err := errors.New("username dan password tidak sesuai")
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	accsesToken, _, err := server.tokenMaker.CreateToken(user.ID, user.Role, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	rsp := models.LoginUserResponse{
		AccsesToken: accsesToken,
		User:        models.NewUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
