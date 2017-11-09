package coms

import "context"

// Communicator provides the Communicate method, which is used to start the
// process of listening for Events on a specific transport method (Such as
// HTTP or XMPP).
type Communicator interface {
	// Communicate sets up a transport method to begin receiving and responding
	// to events.
	//
	// `ctx` can be cancelled to stop the communication process gracefully.
	//
	// `invoker` will be used to run the appropriate handler based on the
	// received Event.
	//
	// An error is returned, or nil if communicator was stopped gracefully.
	Communicate(ctx context.Context, invoker Invoker) error
}
