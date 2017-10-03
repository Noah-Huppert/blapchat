package hndlrs

// Handler is used to process received events.
type Handler interface {
	// Handle will be called once for each event received. The received event
	// will be passed as the `event` argument. An APIResp should be returned, to
	// be sent as a response to the client who sent the event.
	Handle(event Event) APIResp
}

// HandleFn is a type alias for the Handler.Handle() function signature. It is a
// function which takes an Event as its only argument. And returns an APIResp.
type HandleFn func(Event) APIResp

// NewHandler creates a new Handler, given any function of type HandleFn
func NewHandler(hndlFn HandleFn) Handler {
	return FnHandler{hndlFn}
}

// FnHandler implements the Handler interface for any provided function that
// has the HandleFn signature.
type FnHandler struct {
	hndlFn HandleFn
}

func (h FnHandler) Handle(event Event) APIResp {
	return h.hndlFn(event)
}
