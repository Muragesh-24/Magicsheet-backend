package sync

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
)

func RegisterRoutes(api *gin.RouterGroup, handler *Handler) {
	sync := api.Group("/sync")

	protected := sync.Group("/")
	protected.Use(auth.AuthMiddleware())

	protected.POST("/students", middleware.RequireRoles("god", "opc"), handler.SyncStudents)
	protected.POST("/proformas", middleware.RequireRoles("god", "opc"), handler.SyncProformas)
}
