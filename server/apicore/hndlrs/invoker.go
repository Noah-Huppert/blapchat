package hndlrs

// Invoker wraps the Invoke method. Which runs the appropriate resource action
// handler for the provided event.
type Invoker interface {
	// Invoke runs the action specified by the `event` argument.
	//
	// It returns a channel which will receive exactly one APIResp before
	// closing to signify Invoke's completion.
	Invoke(event Event) <-chan APIResp
}
