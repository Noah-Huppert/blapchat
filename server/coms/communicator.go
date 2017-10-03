package coms

import "context"

// Communicators are responsible for interfacing with various transport layers.
// Examples of such transport layers include: HTTP, HTTPS, and XMPP.
//
// Their purpose is to receive client requests from these transports, and translate
// them into events.
//
// These events are then evaluated by an Invoker, and the correct action handler
// is called. The result of this handler is returned by the Invoker. And
// returned as a response by the Communicator.
type Communicator interface {
	// Communicate sets up a transport method to begin receiving and responding
	// to events. It is expected to block execution while it is receiving
	// events.
	//
	// It should use the Invoker provider by the `invoker` argument to the run
	// the appropriate actions when it receives events. Results from the Invoker
	// should then be sent as responses.
	//
	// It returns a channel which will receive exactly 1 `error` to signify Start
	// is complete. This will be `nil` on success. Or an error if one occurred
	// while starting the transport.
	Communicate(invoker Invoker) <-chan error

	// Stop attempts to shut down a transport. So it will no longer send or
	// receive messages.
	//
	// All requests and responses are expected to be done before `ctx` times out.
	//
	// A channel is returned which will receive exactly 1 `error` to signify Stop
	// is complete. This will be `nil` on success. And an error if one occurs
	// while stopping the transport.
	Stop(ctx context.Context) <-chan error
}

// Invokers run the appropriate resource action handler for the given event.
type Invoker interface {
	// Invoke runs the action specified by the `event` argument.
	//
	// It returns a channel which will receive exactly one APIResp before closing
	// to signify Invoke's completion
	Invoke(event Event) <-chan APIResp
}
