package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	config IHTTPConfig
	router *gin.Engine
}

func NewHTTPServer(cnf IHTTPConfig) *HTTPServer {
	return &HTTPServer{
		config: cnf,
		router: gin.Default(),
	}
}

func (s *HTTPServer) GetRouter() *gin.Engine {
	return s.router
}

func (s *HTTPServer) StartHTTPServer() error {
	srv := &http.Server{
		Addr:           s.config.GetAddr(),
		Handler:        s.router,
		ReadTimeout:    time.Duration(s.config.GetReadTimeout()) * time.Second,
		WriteTimeout:   time.Duration(s.config.GetWriteTimeout()) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func Response(ctx *gin.Context, status int, data interface{}) {
	switch v := data.(type) {
	case nil:
		ctx.Status(status)
	default:
		ctx.JSON(status, v)
	}
}

func ErrorResponse(ctx *gin.Context, status int, err error) {
	errs := TransformErrorMessage(err)

	if len(errs) == 0 {
		ctx.AbortWithStatusJSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.AbortWithStatusJSON(status, gin.H{
		"errors": errs,
	})
}
