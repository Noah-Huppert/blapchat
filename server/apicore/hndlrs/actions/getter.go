package actions

// GetAction is the value used throughout code to identify the Get action
const GetAction string = "get"

// Getter wraps a HandleFn, named Get, which should be registered to handle
// Get actions
type Getter interface {
	// Get handles events received by the implementing resource with the
	// Get action
	Get(event Event) APIResp
}

// GetterLoader implements Loader by registering the provided Getter to handle
// Get actions
type GetterLoader struct {
	// Getter is the Getter to register in the Load method
	Getter Getter

	// Target is the Target to register the Getter under
	Target Target
}

// Load implements Loader.Load for GetterLoader
func (g GetterLoader) Load(r Registry) error {
	// Convert to handler
	hndlr := NewHandler(g.Getter.Get)

	// Register
	if err := r.Register(g.Target, hndlr); err != nil {
		return fmt.Errorf("error registering Getter: %s", err.Error())
	}

	// On success, return nil
	return nil
}
