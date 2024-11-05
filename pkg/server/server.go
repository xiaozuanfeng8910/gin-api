package server

import (
	"context"
	"gin-api/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunServer 启动并管理服务器的生命周期
func RunServer(engine *gin.Engine, cfg *config.Config) (*http.Server, error) {
	// 配置服务器
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: engine,
	}

	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()
	log.Printf("服务器已启动, 监听地址 %s", cfg.ServerPort)

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭服务器中...")

	// 设定最大等待时间为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务器优雅关闭失败: %v", err)
		return nil, err
	}

	log.Println("服务器已成功关闭")
	return srv, nil
}
