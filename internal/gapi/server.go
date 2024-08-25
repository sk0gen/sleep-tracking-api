package gapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/pb"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"github.com/sk0gen/sleep-tracking-api/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type Server struct {
	pb.UnimplementedSleepTrackingServer
	logger   *log.Logger
	config   config.Config
	store    db.Store
	jwtMaker *token.JWTMaker
}

func NewServer(cfg config.Config, store db.Store) *Server {

	server := &Server{
		logger:   log.Default(),
		config:   cfg,
		store:    store,
		jwtMaker: token.NewJWTMaker(cfg.AuthConfig.JWTSecret),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("strongpassword", util.StrongPassword)
	}

	return server
}

func (s *Server) Start(ctx context.Context) error {
	cfg := loadConfig()

	grpcServer := grpc.NewServer()
	pb.RegisterSleepTrackingServer(grpcServer, s)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		<-ctx.Done()
		s.logger.Print("grpc: context closed")
		grpcServer.Stop()
		s.logger.Print("grpc: shutting down")
	}()

	s.logger.Printf("grpc: server started on port %s", listener.Addr().String())
	if err = grpcServer.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
