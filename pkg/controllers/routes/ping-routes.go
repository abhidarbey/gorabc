package routes

import (
	"gorabc/pkg/controllers/handlers"

	"github.com/gin-gonic/gin"
)

// Ping Routes function
func Ping(r *gin.Engine) {
	router := r.Group("/")

	router.GET("ping", handlers.Ping)
	router.GET("king", handlers.King)
}
