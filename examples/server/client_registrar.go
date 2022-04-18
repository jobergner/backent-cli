package server

import (
	"sync"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)

// easyjson:skip
type clientRegistrar struct {
	clients         map[*Client]struct{}
	incomingClients map[*Client]struct{}
	mu              sync.Mutex
}

func newClientRegistar() *clientRegistrar {
	return &clientRegistrar{
		clients:         make(map[*Client]struct{}),
		incomingClients: make(map[*Client]struct{}),
	}
}

func (c *clientRegistrar) add(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()

	log.Debug().Str(logging.ClientID, client.id).Msg("adding client")
	c.incomingClients[client] = struct{}{}
}

func (c *clientRegistrar) remove(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()

	log.Debug().Str(logging.ClientID, client.id).Msg("removing client")
	delete(c.clients, client)
	delete(c.incomingClients, client)
}

// TODO unused
func (c *clientRegistrar) kick(client *Client, reason string) {
	log.Debug().Str(logging.ClientID, client.id).Msg("kicking client")
	client.closeConnection(reason)
}

func (c *clientRegistrar) promote(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()

	log.Debug().Str(logging.ClientID, client.id).Msg("promoting client")
	c.clients[client] = struct{}{}
	delete(c.incomingClients, client)
}
