package glesys

import (
	"os"
	"testing"
)

var (
	username string
	apiKey   string
)

func TestMain(m *testing.M) {
	username = os.Getenv("GLESYS_USERNAME")
	apiKey = os.Getenv("GLESYS_API_KEY")

	os.Exit(m.Run())
}
