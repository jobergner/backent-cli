package endtoend

import (
	"context"
	"net/http"
	"time"

	"github.com/jobergner/backent-cli/examples/client"
	"github.com/jobergner/backent-cli/examples/server"
)

const fps = 20

func startServer(m *MockController) func() {
	s := server.NewServer(m, fps)

	signalFail := s.Start()
	go func() {
		err := <-signalFail
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	kill := make(chan struct{})
	go func() {
		<-kill
		s.HttpServer.Shutdown(context.Background())
	}()

	time.Sleep(time.Microsecond * 100)

	return func() {
		kill <- struct{}{}
		time.Sleep(time.Millisecond * 50)
	}
}

func connectClient(m *MockController, onMessageReceived func(b []byte)) (*client.Client, func()) {

	ctx, cancel := context.WithCancel(context.Background())

	c, err := client.NewClient(ctx, m, fps)
	if err != nil {
		panic(err)
	}

	go func() {
		for {

			b := c.ReadUpdate()

			if onMessageReceived != nil {
				onMessageReceived(b)
			}

		}
	}()

	return c, cancel
}
