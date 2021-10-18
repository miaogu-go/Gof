package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/Gof/core"
	"github.com/miaogu-go/Gof/global"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	core.Bootstrap()
	r := gin.New()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, global.GofConfig.System.Mode)
		global.GofLog.Info("test", zap.String("a", "b"))
	})

	r.Run(":8080")
}