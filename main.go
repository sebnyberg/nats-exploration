package main

import (
	"golang.org/x/exp/slog"
)

func main() {
	if err := run(); err != nil {
		slog.Error("err", err)
	}
}

// run initiates a client that re-plays a certain amount of messages from the
// NATS queue in order to catch up to the most recent state.
func run() error {
	natsClient, err := nats.NewClient()
}
