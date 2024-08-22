package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"log"
	"net/http"
)

type Server struct {
	router   *gin.Engine
	logger   *log.Logger
	config   Config
	store    db.Store
	jwtMaker *token.JWTMaker
}

func NewServer(cfg Config, store db.Store) *Server {

	server := &Server{
		logger:   log.Default(),
		config:   cfg,
		store:    store,
		jwtMaker: token.NewJWTMaker(cfg.ApiConfig.JWTSecret),
	}

	server.initRoutes()
	return server
}

func (s *Server) Start() error {
	httpServer := &http.Server{
		Addr:         ":" + s.config.ApiConfig.HttpServerPort,
		Handler:      s.router,
		ReadTimeout:  s.config.ApiConfig.HttpServerTimeout,
		WriteTimeout: s.config.ApiConfig.HttpServerTimeout,
	}

	return httpServer.ListenAndServe()
}

func (s *Server) initRoutes() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.POST("/auth/register", s.CreateUser)
	v1.POST("/auth/login", s.CreateUser)

	v1authRoutes := v1.Group("/").Use(authMiddleware(s.jwtMaker))
	v1authRoutes.POST("/sleep-logs", s.createSleepLog)
	v1authRoutes.GET("/sleep-logs", s.getSleepLogsByUserID)
	s.router = r
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
