package actions

// Getter wraps a HandleFn, named Get, which should be registered to handle
// Get actions
type Getter interface {
	// Get handles events received by the implementing resource with the
	// Get action
	Get(event Event) APIResp
}
