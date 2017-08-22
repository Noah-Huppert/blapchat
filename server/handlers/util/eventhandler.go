package util

import "fmt"

// Return type used by event system
type EventDataType map[string]interface{}

// EventHandler is an interface which defines code to handle specific actions called on resources
type EventHandler interface {
    // GetAction returns the action which the handler will respond to
    GetAction() string

    // GetTransport returns the Transport which the handler will respond to
    GetTransport() Transport

    // Handle is the method which should be called when the server receives an event targeted at the handler
    Handle(data EventDataType) (error, EventDataType)
}
