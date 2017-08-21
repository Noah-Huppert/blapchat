package util

import "fmt"

// EventControllerLoader is an interface for loading EventControllers into the correct backend systems based on transport
type EventControllerLoader interface {
    // GetTransport provides the transport which this EventControllerLoader loads for. Can not be TRANSPORT_BOTH
    GetTransport() *Transport

    // Load registers a specific EventController with the correct backend systems
    Load(c *EventController) error
}

// Loaders holds the EventControllerLoader interfaces that will be used to register new routes with different transports
var Loaders map[string]EventControllerLoader




// ControllerMapType is a type alias for a map, with string keys and EventController interface pointers.
//      The keys: are the ShortString form of an EventTarget, representing the event target resource and action.
type ControllerMapType map[string][]EventController
// Controllers is the variable used to record registered event controllers
var Controllers ControllerMapType

// Register adds the provided EventController to the Controllers map, returns error if provided controller already exists
// in Controllers var.
func Register(c EventController) error {
    // Check if event target exists in map
    short := c.GetEventTarget().ShortString()
    if _, ok := Controllers[short]; !ok {
        // If not, make empty array
        Controllers[short] = []EventController{}
    }

    // Check if controller for specific transport already exists
    for _, tran := range Controllers[short] {
        if tran.GetEventTarget().Transport.Equivalent(c.GetEventTarget().Transport) {
            // If equivalent, then we don't have to add
            return fmt.Errorf("Already registered event controller for event \"%s\" with transport \"%s\"", short, c.GetEventTarget().Transport.Name)
        }
    }

   // If not, add
    Controllers[short] = append(Controllers[short], c)
    return nil
}
