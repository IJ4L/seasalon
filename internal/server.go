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
	produk     repository.Produk
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, produk repository.Produk) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{produk: produk, config: config, tokenMaker: token}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))

	authRoutes.POST("/produk", server.CreateProduk)
	authRoutes.GET("/produk", server.GetAllProduk)
	authRoutes.DELETE("/produk", server.DestroyProduk)
	authRoutes.PATCH("/produk", server.UpdateProduk)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}