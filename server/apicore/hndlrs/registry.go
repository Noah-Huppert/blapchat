package hndlrs

// Registry provides an interface to add and query event action handlers
type Registry interface {
	// Register adds the Handler provided in the `hndlr` argument to the central
	// registry under the provided Target. The specifics of this registry
	// are left up to the implementation.
	Register(t Target, hndlr Handler) error

	// WithTarget returns the Handler which is registered under the Target
	// provided by the `t` argument. An error is returned if the one occurs,
	// otherwise nil will be returned.
	WithTarget(t Target) (Handler, error)
}
