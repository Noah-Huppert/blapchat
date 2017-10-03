package hndlrs

// Loader provides an interface for registering a group of handlers in a
// Registry.
type Loader interface {
	// Load should register any required Handlers with the Registry provided in
	// the `r` argument. An error should be returned if one occurs, otherwise
	// `nil`.
	Load(r Registry) error
}
