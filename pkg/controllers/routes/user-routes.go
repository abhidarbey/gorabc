package routes

import (
	"gorabc/pkg/controllers/handlers"

	"github.com/gin-gonic/gin"
)

// Users Routes function
func Users(r *gin.Engine) {
	h := handlers.UserHandler

	router := r.Group("api/users")

	router.POST("", h.Create)
	router.GET("", h.FindAll)
	router.GET(":id", h.GetByID)
	router.PUT(":id", h.Update)
	router.PUT(":id/password", h.UpdatePassword)
	router.DELETE(":id", h.Delete)
}
