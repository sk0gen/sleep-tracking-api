package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/util"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	cfg := &Config{
		ApiConfig: ApiConfig{
			JWTSecret:          util.RandomString(32),
			JWTTokenExpiration: time.Minute,
		},
	}

	server := NewServer(*cfg, store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
