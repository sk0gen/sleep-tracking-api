package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	HttpServerPort    string        `env:"HTTP_SERVER_PORT,required"`
	HttpServerTimeout time.Duration `env:"HTTP_SERVER_TIMEOUT" ,envDefault:"10s"`
	JWTSecret         string        `env:"JWT_SECRET"`
}

type Server struct {
	router   *gin.Engine
	logger   *log.Logger
	config   *Config
	database db.Store
}

func NewServer(database db.Store) *Server {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}

	server := &Server{
		logger:   log.Default(),
		config:   &cfg,
		database: database,
	}

	server.initRoutes()
	return server
}

func (s *Server) Start() error {
	httpServer := &http.Server{
		Addr:         ":" + s.config.HttpServerPort,
		Handler:      s.router,
		ReadTimeout:  s.config.HttpServerTimeout,
		WriteTimeout: s.config.HttpServerTimeout,
	}

	return httpServer.ListenAndServe()
}

func (s *Server) initRoutes() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	v1.POST("/sleep-logs", s.createSleepLog)
	v1.GET("/sleep-logs", s.getSleepLogsByUserID)
	s.router = r
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
