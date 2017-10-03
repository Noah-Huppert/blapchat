package actions

// Updater wraps a HandleFn, named Update, which should be registered to handle
// the Update action
type Updater interface {
	// Update handles events received by the implementing resource with the
	// Update action
	Update(event Event) APIResp
}
