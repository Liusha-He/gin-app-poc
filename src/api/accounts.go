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

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// Create Account 	godoc
// @Summary 		create account
// @Schemes 		http
// @Description 	Takes an account json and store in DB, Returned saved json.
// @Tags 			accounts
// @Produce 		json
// @Param 			account  body	createAccountRequest true  "account json"
// @Success 		200 {object} dao.Account
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

// GetAccountByID 	godoc
// @Summary 		get account by id
// @Schemes 		http
// @Description 	Takes an account id with path, Returned account info json.
// @Tags 			accounts
// @Produce 		json
// @Param 			id path string true "search by id"
// @Success 		200	{object} dao.Account
// @Router 			/api/v1/accounts/{id} [get]
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary 		Get account list
// @Schemes 		http
// @Description		Responds with the list of accounts
// @Tags 			accounts
// @Produce 		json
// @Success 		200	{array} dao.Account
// @Router 			/api/v1/accounts [get]
func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := dao.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
