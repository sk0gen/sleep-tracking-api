package db

import (
	"context"
	"os"
	"testing"
)

var testStore Store
var testDatabase *TestDatabase

func TestMain(m *testing.M) {
	testDatabase = NewTestDatabase()

	testStore = NewStore(context.Background(), testDatabase.Config)
	code := m.Run()
	close()
	os.Exit(code)
}

func close() {
	testStore.Close()
	testDatabase.Close()
}
