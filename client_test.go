package glesys

import (
	"os"
	"testing"
)

var (
	username string
	apiKey   string
	c        *Client
)

func TestMain(m *testing.M) {
	username = os.Getenv("GLESYS_USERNAME")
	apiKey = os.Getenv("GLESYS_API_KEY")

	c = NewClient(username, apiKey, nil)

	os.Exit(m.Run())
}
