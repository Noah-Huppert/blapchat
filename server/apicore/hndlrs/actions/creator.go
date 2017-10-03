package actions

// Creator wraps a HandleFn, named Create, which should be registered to handle
// the Create action.
type Creator interface {
	// Create handles events received by the implementing resource with the
	// Create action
	Create(event Event) APIResp
}
