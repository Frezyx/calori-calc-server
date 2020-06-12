package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=123456 dbname=calori-calc sslmode=disable"
	}
	os.Exit(m.Run())
}
