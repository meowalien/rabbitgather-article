package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/meowalien/rabbitgather-article/conf"
	"github.com/meowalien/rabbitgather-article/global"
	"github.com/meowalien/rabbitgather-article/server/handler"
	"github.com/meowalien/rabbitgather-lib/wrapper"
	"net/http"
)

type Server struct {
	Debug  bool
	Config conf.RestfulServerConfiguration

	serverInst              *http.Server
	ginEngine               *wrapper.GinEngine
	shutdownCallbackMethods []func() error
}

func (w *Server) Start(ctx context.Context) {
	global.Logger.Println("APIServer listen on : ", w.Config.Port)

	w.ginEngine = &wrapper.GinEngine{
		Engine:       gin.Default(),
	}
	serverInst := &http.Server{
		Addr:    ":" + w.Config.Port,
		Handler: w.ginEngine,
	}
	w.MountService(ctx)
	go func() {
		if err := serverInst.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("error when start Server", err)
		}
	}()
	global.Logger.Info("APIServer Started .")
}

func (h *Server) Stop(ctx context.Context) error {
	return nil
}

const frontEndWebURL = "https://meowalien.com:443"

func (w *Server) MountService(ctx context.Context) {
	// 下屬Group都會繼承 AllowOrigins 屬性
	w.ginEngine.AllowOrigins =  []string{frontEndWebURL}

	articleGroup := w.ginEngine.Group("/article")
	{
		basicHandler := handler.Basic{}
		// 基本文章資訊
		articleGroup.GET("/basic", basicHandler.Get)
		articleGroup.POST("/basic", basicHandler.POST)
		articleGroup.DELETE("/basic", basicHandler.DELETE)
		articleGroup.PATCH("/basic", basicHandler.PATCH)
	}

	{
		listenGroup := articleGroup.Group("/listen")
		// websocket 監聽文章變更
		listenGroup.GET("/")
	}
}
