package server

import (
	"fmt"
	"gitlab/go-prolog-api/example/repository"
	"gitlab/go-prolog-api/example/security/middleware"
	"gitlab/go-prolog-api/example/security/token"
	"gitlab/go-prolog-api/example/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	repo       repository.Repo
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, repo repository.Repo) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{repo: repo, config: config, tokenMaker: token}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	
	router.POST("/user/register", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))

	authRoutes.POST("/service", server.CreateService)
	authRoutes.GET("/services", server.GetAllService)
	authRoutes.GET("/service/:service_id", server.GetService)
	authRoutes.DELETE("/service/:service_id", server.DestroyService)
	authRoutes.PUT("/service", server.UpdateService)

	authRoutes.POST("/branch", server.CreateBranch)
	authRoutes.GET("/branchs", server.GetBranchs)
	authRoutes.GET("/branch/:branch_id", server.GetBranch)
	authRoutes.DELETE("/branch/:branch_id", server.DestroyBranch)
	authRoutes.PUT("/branch", server.UpdateBranch)

	authRoutes.POST("/reservation", server.CreateReservation)
	authRoutes.GET("/reservations", server.GetReservations)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
