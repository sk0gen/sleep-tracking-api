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

	var err error
	testStore, err = NewStore(context.Background(), testDatabase.Config)
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
