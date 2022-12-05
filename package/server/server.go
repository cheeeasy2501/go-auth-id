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

func (s *HTTPServer) StartHTTPServer() {
	srv := &http.Server{
		Addr:           ":9090",
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv.ListenAndServe()
}
