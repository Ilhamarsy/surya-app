package api

import (
	"net/http"
	db "projek-abal-abal/db/sqlc"
	"projek-abal-abal/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	publicRoute := router.Group("/")
	authRoute := router.Group("/").Use(middleware.DeserializeUser(store))

	authRoute.GET("/me", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("currentUser").(db.User)
		ctx.JSON(http.StatusOK, currentUser)
	})
	publicRoute.POST("/register", server.createUser)
	publicRoute.POST("/login", server.loginUser)

	authRoute.PUT("/refresh", server.refreshAccessToken)

	publicRoute.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"}) })
	// r.HandleFunc("/register", usercontroller.Register).Methods("POST")
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	helper.ResponseJson(w, http.StatusOK, map[string]string{"message": "Hello World"})
	// }).Methods("GET")

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
