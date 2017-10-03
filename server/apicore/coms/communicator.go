package coms

import "context"

// Communicator provides the Communicate method, which is used to start the
// process of listening for Events on a specific transport method (Such as
// HTTP or XMPP).
type Communicator interface {
	// Communicate sets up a transport method to begin receiving and responding
	// to events. It is expected to block execution while it is receiving
	// events.
	//
	// It should use the Invoker provider by the `invoker` argument to the run
	// the appropriate actions when it receives events. Results from the Invoker
	// should then be sent as responses.
	//
	// It returns a channel which will receive exactly 1 `error` to signify
	// Start is complete. This will be `nil` on success. Or the error if
	// one occurred.
	Communicate(invoker Invoker) <-chan error

	// Stop attempts to shut down a transport. So it will no longer send or
	// receive messages.
	//
	// All requests and responses are expected to be done before `ctx`
	// times out.
	//
	// A channel is returned which will receive exactly 1 `error` to
	// signify Stop is complete. This will be `nil` on success. And an
	// error if one occurs.
	Stop(ctx context.Context) <-chan error
}
