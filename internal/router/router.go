package router

import (
	"gin-api/pkg/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitRoutes 初始化路由
func InitRoutes(logger *zap.Logger, db *gorm.DB, validator validation.Validator) *gin.Engine {
	r := gin.New() // 使用 gin.New() 创建一个新的实例
	// 创建 /api 路由组
	apiGroup := r.Group("/api")
	InitUserRoutes(apiGroup.Group("/user"), logger, db, validator)
	return r
}
