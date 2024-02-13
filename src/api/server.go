package api

import (
	"fmt"
	"simple-bank/src/auth"
	"simple-bank/src/dao"
	"time"

	docs "simple-bank/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}

type Server struct {
	store      dao.Store
	tokenMaker auth.Maker
	Router     *gin.Engine
	config     Config
}

// NewServer will create the http Server
func NewServer(config Config, store dao.Store) (*Server, error) {
	tokenMaker, err := auth.NewPASETOMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("unable to generate token: %s", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	docs.SwaggerInfo.BasePath = "/"

	// swager router
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")

	// Register all routers below
	server.UsersRouter(v1)
	server.AccountsRouters(v1)
	server.TransfersRouters(v1)

	server.Router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
