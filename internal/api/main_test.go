package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"github.com/sk0gen/sleep-tracking-api/util"
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
			JWTSecret:          util.RandomString(32),
			JWTTokenExpiration: time.Minute,
		},
	}

	server := NewServer(*cfg, testStore)

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
