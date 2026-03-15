package api

import (
	"fmt"

	sqlc "auth-system/internal/db/sqlc"
	"auth-system/internal/token"
	"auth-system/internal/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *sqlc.Queries
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *sqlc.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PasetoKey)
	if err != nil {
		return nil, fmt.Errorf("token maker oluşturulamadı: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/register", server.registerUser)
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/anasayfa", server.homePage)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
