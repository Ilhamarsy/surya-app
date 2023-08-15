package middleware

import (
	"context"
	"fmt"
	"net/http"
	db "projek-abal-abal/db/sqlc"
	"projek-abal-abal/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(db *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := util.LoadConfig(".")
		sub, err := util.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		id, _ := strconv.ParseInt(fmt.Sprint(sub), 10, 64)

		user, err := db.GetUserById(context.TODO(), id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
