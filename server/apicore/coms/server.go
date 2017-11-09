package coms

import (
	"context"
)

// Server ties together Communicators, Invokers, and Handlers to respond to 
// client requests from multiple message transport mechanisms.
type Server interface {
	// Communicator is the interface to use for translating raw received 
	// messages into Events.
	//
	// This allows the same handler logic to be used regardless of the 
	// message transport method.
	Communicator Communicator

	// Invoker is the interface used to run the correct handler for  
	// received events.
	Invoker Invoker

	// Registry is the interface used to store which handlers respond to 
	// which events
	Registry Registry
}

// Start begins the process of listening for and responding to requests.
//
// `ctx` will be used to derive all request Contexts. And can be cancelled to 
// stop server execution.
//
// It returns an error, or nil if the server is stopped gracefully. 
func (s Server) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			s.Communicate(ctx, s.Invoker)
		}
	}

	return eChan
}
