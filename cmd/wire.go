package main

import (
	"gin-api/internal/config"
	"gin-api/internal/router"
	"gin-api/pkg/db"
	"gin-api/pkg/log"
	"gin-api/pkg/server"
	"gin-api/pkg/validation"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	server *http.Server
	db     *gorm.DB
	logger *zap.Logger
}

// NewApp 是 App 的构造函数，由 wire 调用
func NewApp(server *http.Server, db *gorm.DB, logger *zap.Logger) *App {
	return &App{
		server: server,
		db:     db,
		logger: logger,
	}
}

// Run 启动应用服务器
func (a *App) Run() error {
	return a.server.ListenAndServe()
}

// Shutdown 优雅关闭服务器
func (a *App) Shutdown() {
	if err := a.server.Close(); err != nil {
		a.logger.Error("服务器关闭失败", zap.Error(err))
	}
	if sqlDB, err := a.db.DB(); err == nil {
		sqlDB.Close()
	}
}

func InitializeApp() (*App, error) {
	wire.Build(
		config.InitializeConfig, // 初始化配置
		log.InitializerLog,      // 初始化日志
		db.InitMySQLGorm,        // 初始化数据库
		wire.NewSet(
			validation.NewCustomValidator,
			wire.Bind(new(validation.Validator), new(*validation.CustomValidator)),
		), // 统一校验器
		router.InitRoutes, // 初始化路由
		server.RunServer,  // 启动服务器
		NewApp,
	)
	return &App{}, nil
}
