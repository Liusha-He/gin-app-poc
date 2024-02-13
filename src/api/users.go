package api

import (
	"database/sql"
	"net/http"
	"simple-bank/src/dao"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func NewUserResponse(user dao.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
		Email:             user.Email,
	}
}

// CreateUser 	godoc
// @Summary 		create user
// @Schemes 		http
// @Description 	Takes an user json and store in DB, Returned saved json.
// @Tags 			users
// @Produce 		json
// @Param 			user  body	createUserRequest true  "user json"
// @Success 		200 {object} userResponse
// @Router 			/api/v1/users [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := dao.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := dao.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := NewUserResponse(user)
	ctx.JSON(http.StatusCreated, resp)
}

// Login 	godoc
// @Summary 		login
// @Schemes 		http
// @Description 	Takes an login user json and store in DB, Returned saved json.
// @Tags 			users
// @Produce 		json
// @Param 			login  body	loginRequest true  "login json"
// @Success 		200 {object} loginResponse
// @Router 			/api/v1/users/login [post]
func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// check if the provided password match the user
	err = dao.CheckPassword(user.HashedPassword, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// if no error, provide the user the access token
	token, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginResponse{
		AccessToken: token,
		User:        NewUserResponse(user),
	}
	ctx.JSON(http.StatusOK, resp)
}
