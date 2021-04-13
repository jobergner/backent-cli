package main

import (
	server "bar-cli/serverfactory/server_example/server"
)

func main() {
	server.Start(actions{
		movePlayer: func(a PlayerID, x float64, y float64, e *Engine) {
			log.Println("moving player..")
			e.Player(a).Position(e).SetX(e, x)
		},
	})
}
