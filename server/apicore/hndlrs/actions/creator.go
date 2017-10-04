package actions

// CreateAction is the value used throughout code to identify the Create action
const CreateAction string = "create"

// Creator wraps a HandleFn, named Create, which should be registered to handle
// the Create action.
type Creator interface {
	// Create handles events received by the implementing resource with the
	// Create action
	Create(event Event) APIResp
}

// CreatorLoader implements Loader by registering a provided Creator interface
// to handle Create actions
type CreatorLoader struct {
	// Creator is the Creator to register in the Load method
	Creator Creator

	// Target is the Target at which the Creator will be registered
	Target Target
}

// Load implements Loader.Load from CreatorLoader
func (c CreatorLoader) Load(r Registry) error {
	// Convert c.Creator into a Handler
	hndlr := NewHandler(c.Creator.Create)

	// Attempt to register
	if err := r.Register(c.Target, hndlr); err != nil {
		return fmt.Errorf("error registering Creator: %s", err.Error())
	}

	// On success, return nil
	return nil
}
