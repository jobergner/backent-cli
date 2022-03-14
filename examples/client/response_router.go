package client

import "sync"

// easyjson:skip
type responseRouter struct {
	pending map[string]chan []byte
	mu      sync.Mutex
}

func (r *responseRouter) add(id string, ch chan []byte) {
	r.mu.Lock()
	r.pending[id] = ch
	r.mu.Unlock()
}

func (r *responseRouter) remove(id string) {
	r.mu.Lock()
	ch := r.pending[id]
	delete(r.pending, id)
	close(ch)
	r.mu.Unlock()
}

func (r *responseRouter) route(msg Message) {
	ch, ok := r.pending[msg.ID]

	if !ok {
		return
	}

	ch <- msg.Content
}
