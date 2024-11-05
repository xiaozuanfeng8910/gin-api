package router

import (
	"gin-api/internal/handlers"
	"gin-api/pkg/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitUserRoutes 初始化用户路由
func InitUserRoutes(userGroup *gin.RouterGroup, logger *zap.Logger, db *gorm.DB, validator validation.Validator) {
	userHandler := handlers.NewUserHandler(logger, db, validator)
	userGroup.GET("/list", userHandler.GetUsers)
	userGroup.GET("/detail/:id", userHandler.GetUserInfo)
	userGroup.POST("/create", userHandler.CreateUser)
	userGroup.PUT("/update/:id", userHandler.UpdateUser)
	userGroup.DELETE("/delete/:id", userHandler.DeleteUser)
}
