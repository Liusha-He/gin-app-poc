package api

import "github.com/gin-gonic/gin"

func (server *Server) AccountsRouters(r *gin.RouterGroup) {
	r.POST("/accounts", server.createAccount)
}