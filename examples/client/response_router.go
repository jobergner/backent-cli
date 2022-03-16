package client

import (
	"sync"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)

// easyjson:skip
type responseRouter struct {
	pending map[string]chan []byte
	mu      sync.Mutex
}

func (r *responseRouter) add(id string, ch chan []byte) {
	r.mu.Lock()
	log.Debug().Str(logging.MessageID, id).Msg("adding channel to router")
	r.pending[id] = ch
	r.mu.Unlock()
}

func (r *responseRouter) remove(id string) {
	r.mu.Lock()
	log.Debug().Str(logging.MessageID, id).Msg("removing channel to router")
	ch := r.pending[id]
	delete(r.pending, id)
	close(ch)
	r.mu.Unlock()
}

func (r *responseRouter) route(response Message) {
	ch, ok := r.pending[response.ID]

	if !ok {
		log.Warn().Str(logging.MessageID, response.ID).Msg("cannot find channel for routing response")
		return
	}

	log.Debug().Str(logging.MessageID, response.ID).Msg("routing response")
	ch <- response.Content
}
