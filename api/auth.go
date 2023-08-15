package api

import (
	"fmt"
	"net/http"
	db "projek-abal-abal/db/sqlc"
	"projek-abal-abal/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type refreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (server *Server) refreshAccessToken(ctx *gin.Context) {
	var access_token_auth string
	var req refreshAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	config, _ := util.LoadConfig(".")
	sub, err := util.ValidateToken(req.RefreshToken, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		access_token_auth = fields[1]
	}

	currentUser := ctx.MustGet("currentUser").(db.User)

	if err := util.ValidateAccessWithRefresh(access_token_auth, req.RefreshToken); err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(fmt.Sprint(sub), 10, 64)

	if currentUser.ID != id {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail"})
		return
	}

	user, err := server.store.GetUserById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := util.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := util.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token, "refresh_token": refresh_token})
}
