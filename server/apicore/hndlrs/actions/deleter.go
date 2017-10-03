package actions

// Deleter wraps a HandleFn, named Delete, which should be registered to handle
// the Delete action
type Deleter interface {
	// Delete handles any events received by the implementing resource with
	// the Delete action
	Delete(event Event) APIResp
}
