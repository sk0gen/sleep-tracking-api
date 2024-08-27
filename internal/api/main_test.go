package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"go.uber.org/zap"
	"os"
	"testing"
	"time"
)

var testStore db.Store
var testDatabase *db.TestDatabase

func newTestServer(t *testing.T) *Server {
	t.Helper()

	cfg := &config.Config{
		Database: testDatabase.Config,
		AuthConfig: token.Config{
			JWTSecret:          "secret",
			JWTTokenExpiration: time.Hour,
		},
	}

	logger, _ := zap.NewDevelopment()
	server := NewServer(*cfg, testStore, logger)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	testDatabase = db.NewTestDatabase()

	var err error
	testStore, err = db.NewStore(context.Background(), testDatabase.Config)
	if err != nil {
		close()
	}
	code := m.Run()
	close()
	os.Exit(code)
}

func close() {
	testStore.Close()
	testDatabase.Close()
}
