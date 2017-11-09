package hndlrs

// Invoker wraps the Invoke method. Which runs the appropriate resource action
// handler for the provided event.
type Invoker interface {
	// Invoke runs the action specified by the `event` argument.
	//
	// `ctx` can be cancelled to stop execution of a handler.
	//
	// It returns a channel which will receive exactly one error before
	// closing, to signify Invoke's completion. Nil on success.
	Invoke(ctx context.Context, event Event) <-chan error
}
