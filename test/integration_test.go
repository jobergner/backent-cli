package test

import (
	"testing"

	"github.com/jobergner/backent-cli/examples/client"
	"github.com/jobergner/backent-cli/examples/server"
)

func TestIntegration(t *testing.T) {

	server.Start()

	_ = client.NewClient()
}
