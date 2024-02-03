package api

import (
	"net/http"
	"simple-bank/src/dao"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR GBP CNY"`
}

// Create Account 	godoc
// @Summary 		create account
// @Schemes 		http
// @Description 	Takes an account json and store in DB, Returned saved json.
// @Tags 			accounts
// @Produce 		json
// @Param 			account  body	createAccountRequest true  "account json"
// @Success 		200
// @Router 			/api/v1/accounts [post]
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := dao.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0.00,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}
