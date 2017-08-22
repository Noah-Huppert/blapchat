package util

import "fmt"

// HandlerLoader is an interface for loading EventControllers into the correct backend systems based on transport
type HandlerLoader interface {
    // GetTransport provides the transport which this HandlerLoader loads for. Can not be TRANSPORT_BOTH
    GetTransport() *Transport

    // Load registers a specific EventHandler with the correct backend systems
    Load(h *EventHandler) error
}

// Loaders holds the HandlerLoader interfaces that will be used to register new routes with different transports
var Loaders map[string]HandlerLoader

// Handlers holds all EventHandlers application responds to. Keys are:
//      Handlers[EventHandler Resource][EventHandler Action]
var Handlers map[string]map[string][]*EventHandler

// AddHandler adds the provided EventHandler to the Handlers map, under the specified resource name
func AddHandler(resource string, h *EventHandler) error {
    // Make map for resource if doesnt't exist
    if _, ok := Handlers[resource]; !ok {
        Handlers[resource] = map[string][]*EventHandler{}
    }

    resMap := Handlers[resource]

    // Make array for action if doesn't exist
    if _, ok := resMap[h.GetAction()]; !ok {
        resMap[h.GetAction()] = []*EventHandler{}
    }

    actMap := resMap[h.GetAction()]

    // Check if equivalent
}
