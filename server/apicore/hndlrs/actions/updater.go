package actions

// UpdateAction is the value used throughout code to identify the Update action
const UpdateAction string = "update"

// Updater wraps a HandleFn, named Update, which should be registered to handle
// the Update action
type Updater interface {
	// Update handles events received by the implementing resource with the
	// Update action
	Update(event Event) APIResp
}

// UpdaterLoader implements Loader by register the provided Updater to handle
// Update actions
type UpdaterLoader struct {
	// Updater is the Updater the Load method will register
	Updater Updater

	// Target is the Target the Updater will be register under
	Target Target
}

// Load implements Loader.Load for UpdaterLoader
func (u UpdaterLoader) Load(r Registry) error {
	// Convert to handler
	hndlr := NewHandler(u.Updater.Update)

	// Register
	if err := r.Register(u.Target, hndlr); err != nil {
		return fmt.Errorf("error registering Updater: %s", err.Error())
	}

	// On success, return nil
}
