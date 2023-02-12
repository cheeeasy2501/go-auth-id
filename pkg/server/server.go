package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router *gin.Engine
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: gin.Default(),
	}
}

func (s *HTTPServer) GetRouter() *gin.Engine {
	return s.router
}

func (s *HTTPServer) StartHTTPServer() error {
	srv := &http.Server{
		Addr:           ":9090",
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
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
		ctx.AbortWithStatusJSON(status,  gin.H{
			"error": err.Error(),
		})
		return
	} 

	ctx.AbortWithStatusJSON(status, gin.H{
		"errors": errs,
	})
}
