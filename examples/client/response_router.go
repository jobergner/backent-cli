package client

import (
	"sync"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
)

func newReponseRouter() *responseRouter {
	return &responseRouter{
		pending: make(map[int]chan []byte),
	}
}

// easyjson:skip
type responseRouter struct {
	pending map[int]chan []byte
	mu      sync.Mutex
}

func (r *responseRouter) add(id int, ch chan []byte) {
	r.mu.Lock()

	r.pending[id] = ch

	r.mu.Unlock()
}

func (r *responseRouter) remove(id int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ch, ok := r.pending[id]
	if !ok {
		return
	}

	delete(r.pending, id)
	close(ch)
}

func (r *responseRouter) route(response Message) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ch, ok := r.pending[response.ID]
	if !ok {
		log.Warn().Int(logging.MessageID, response.ID).Msg("cannot find channel for routing response")
		return
	}

	ch <- response.Content
}
