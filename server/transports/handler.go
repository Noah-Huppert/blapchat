package transports

import (
    "fmt"
)

type ActionType string
const (
    Create ActionType = "create"
    Get    ActionType = "get"
    Update ActionType = "update"
    Delete ActionType = "delete"
)

// HandlerTarget holds data which identifies the types of requests a Handler will Handle
type HandlerTarget struct {
    // Path is a list of strings which will be converted into a location identifier by a transport. The way this is done
    // depends on the transport.
    //
    // For example the HttpTransport concatenates Path items with slashes to make a regular Url path
    Path []string

    // Action is the ActionType which the handler will handle requests for
    Action ActionType

    // TransportRestrictions is a map with TransportTypes keys. Boolean values specify if the HandlerTarget includes the
    // TransportType specified by the key. The true value marks a TransportType as included, a false value: excluded.
    //
    // If this map is empty then all transport types are allowed
    TransportRestrictions map[TransportType]bool
}

// StandardHandlerTarget creates a new HandlerTarget instance with the specified Path and Action. Along with an empty
// TransportRestrictions map
func StandardHandlerTarget(path []string, action ActionType) HandlerTarget {
    return HandlerTarget{path, action, map[TransportType]bool{}}
}

// String converts a HandlerTarget into a string
func (t HandlerTarget) String() string {
    return fmt.Sprint(t)
}

// IncludesTransport returns true if the provided transport is included in the HandlerTarget, via the
// .TransportRestrictions map.
func (t HandlerTarget) IncludesTransport(transport TransportType) bool {
    // Check if any transport restrictions are specified
    if len(t.TransportRestrictions) == 0 {
        // If no restrictions, allow transport
        return true
    }

    // Otherwise check if transport exists in map
    if allowed, exists := t.TransportRestrictions[transport]; exists {
        // If exists, return allowed value
        return allowed
    }

	// If doesn't exist, not allowed
    return false
}

// HandlerData is a type alias for a map with string keys and any values. It is used to pass data from requests or for
// responses
type HandlerData map[string]interface{}

// Handler interface defines the Handle method. Which receives data from a client request and responds with a data
// payload or an error
type Handler interface {
	// Target returns a HandlerTarget struct containing all the information required to determine which requests this
    // handlers will respond to
    Target() HandlerTarget

    // Handle recieves client request data and returns a data payload or an error. A data payload and an error can not
    // be returned at the same time
    Handle(data HandlerData) (error, HandlerData)
}
