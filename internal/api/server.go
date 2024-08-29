package api

import (
	"context"
	"errors"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	docs "github.com/sk0gen/sleep-tracking-api/internal/api/docs"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	"github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"github.com/sk0gen/sleep-tracking-api/util"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// @title Sleep Tracking Api
// @version 0.1
// @description

// @contact.name Wojciech Gawinski
// @contact.url https://github.com/sk0gen/sleep-tracking-api

// @license.name MIT License
// @license.url https://github.com/stefanprodan/podinfo/blob/master/LICENSE

// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" and then your API Token

type Server struct {
	router    *gin.Engine
	logger    *zap.Logger
	apiConfig Config
	config    config.Config
	store     db.Store
	jwtMaker  *token.JWTMaker
}

func NewServer(cfg config.Config, store db.Store, logger *zap.Logger) *Server {
	apiCfg := loadConfig()

	server := &Server{
		logger:    logger,
		config:    cfg,
		apiConfig: apiCfg,
		store:     store,
		jwtMaker:  token.NewJWTMaker(cfg.AuthConfig.JWTSecret),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("strongpassword", util.StrongPassword)
	}

	server.initRoutes()
	return server
}

func (s *Server) Start(ctx context.Context) error {

	srv := &http.Server{
		Addr:         ":" + s.apiConfig.HttpServerPort,
		Handler:      s.router,
		ReadTimeout:  s.apiConfig.HttpServerTimeout,
		WriteTimeout: s.apiConfig.HttpServerTimeout,
	}

	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()

		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()
		s.logger.Info("server.Serve: shutting down")
		errCh <- srv.Shutdown(shutdownCtx)
	}()

	s.logger.Info("REST: Server started on port", zap.String("port", srv.Addr))
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("server.ListenAndServe", zap.Error(err))
		return err
	}

	// Return any errors that happened during shutdown.
	if err := <-errCh; err != nil {
		s.logger.Error("server.Shutdown", zap.Error(err))
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	return nil
}

func (s *Server) initRoutes() {
	r := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.Use(ginzap.Ginzap(s.logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(s.logger, true))

	v1 := r.Group("/api/v1")
	v1.POST("/auth/register", s.createUser)
	v1.POST("/auth/login", s.loginUser)

	v1authRoutes := v1.Group("/").Use(authMiddleware(s.jwtMaker))
	v1authRoutes.POST("/sleep-logs", s.createSleepLog)
	v1authRoutes.GET("/sleep-logs", s.getSleepLogs)
	v1authRoutes.DELETE("/sleep-logs/:id", s.deleteSleepLogByID)
	v1authRoutes.PUT("/sleep-logs/:id", s.updateSleepLogByID)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.router = r
}

type errResponse struct {
	Error string `json:"error"`
}

func errorResponse(err error) errResponse {
	return errResponse{Error: err.Error()}
}
