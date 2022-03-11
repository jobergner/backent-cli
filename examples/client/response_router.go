package state

type responseRouter struct {
	pending map[string]chan []byte
}

func (r *responseRouter) add(id string, ch chan []byte) {
	r.pending[id] = ch
}

func (r *responseRouter) remove(id string) {
	ch := r.pending[id]
	delete(r.pending, id)
	close(ch)
}

func (r *responseRouter) route(id string, contentBytes []byte) {
	ch, ok := r.pending[id]

	if !ok {
		return
	}

	ch <- contentBytes
}
