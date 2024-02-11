package api

import (
	"simple-bank/src/dao"

	docs "simple-bank/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	store  dao.Store
	Router *gin.Engine
}

// NewServer will create the http Server
func NewServer(store dao.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	// swager router
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")

	// Register all routers below
	server.AccountsRouters(v1)

	server.Router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
