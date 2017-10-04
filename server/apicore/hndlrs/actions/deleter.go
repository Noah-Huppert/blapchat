package actions

// DeleteAction is the value used throughout code to identify the Delete action
const DeleteAction string = "delete"

// Deleter wraps a HandleFn, named Delete, which should be registered to handle
// the Delete action
type Deleter interface {
	// Delete handles any events received by the implementing resource with
	// the Delete action
	Delete(event Event) APIResp
}

// DeleterLoader implements the Loader interface by registering the provided
// Deleter to handle delete actions
type DeleterLoader struct {
	// Deleter is the Deleter registered in the Load method
	Deleter Deleter

	// Target is the Target which the Deleter will be registered for
	Target Target
}

// Load implements Loader.Load for DeleterLoader
func (d DeleterLoader) Load(r Registry) error {
	// Convert to handler
	hdnlr := NewHandler(d.Deleter.Delete)

	// Register
	if err := r.Register(d.Target, hndlr); err != nil {
		return fmt.Errorf("error registering Deleter: %s", err.Error())
	}

	// On success, return nil
	return nil
}
