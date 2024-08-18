package api

import (
	"github.com/gin-gonic/gin"
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
	router *gin.Engine
	logger *log.Logger
	config *Config
}

func Run() {

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}
	routersInit := initRoutes()

	server := &http.Server{
		Addr:         ":" + cfg.HttpServerPort,
		Handler:      routersInit,
		ReadTimeout:  cfg.HttpServerTimeout,
		WriteTimeout: cfg.HttpServerTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server start failed: %s", err)
	}
}

func initRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	return r
}
