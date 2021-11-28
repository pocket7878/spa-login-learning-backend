package presentation

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pocket7878/spa_login_learning_backend/infrastructure/middleware"
)

// New sets up our routes and returns a *gin.Engine.
func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"}, //TODO allow from Vercel hosted frontend
			AllowCredentials: true,
			AllowHeaders:     []string{"Authorization"},
		},
	))

	// Public endpoints
	router.Any("/", func(ctx *gin.Context) {
		response := map[string]string{
			"message": "Hello, World!",
		}
		ctx.JSON(http.StatusOK, response)
	})

	router.Any("/greeting", func(ctx *gin.Context) {
		response := map[string]string{
			"message": "Howdy?",
		}
		ctx.JSON(http.StatusOK, response)
	})

	// Require authz
	// TODO: Retrieve user real todos from usecase.
	router.GET(
		"/todos",
		middleware.EnsureValidToken(),
		func(ctx *gin.Context) {
			response := []map[string]string{
				{
					"description": "This is todo 1",
				},
				{
					"description": "This is todo 2",
				},
			}
			ctx.JSON(http.StatusOK, response)
		},
	)

	return router
}
