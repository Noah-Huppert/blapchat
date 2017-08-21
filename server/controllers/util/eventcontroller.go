package util

import "fmt"

// Return type used by event system
type EventReturnType map[string]interface{}

// EventTarget is a struct which holds the information needed to determine an event's target. This is the target resource,
// action and the transportation method.
type EventTarget struct {
    Resource string
    Action string
    Transport *Transport
}

// ShortString returns the shorter string representation of an event target, the form: {{ .Resource }}.{{ .Action }}
func (e *EventTarget) ShortString() string {
    return fmt.Sprintf("%s.%s", e.Resource, e.Action)
}

// Interface which provides event handler and data about target event
type EventController interface {
    // GetEventTarget returns a struct which provides information on the target event for this controller
    GetEventTarget() *EventTarget

    // Handle fn to run when event is received
    Handle(data EventReturnType) (error, EventReturnType)
}

// Default implementation of EventController which provides event target data from Resource and Action member variables
type DefaultEventController struct {
    // Event target data
    Resource string
    Action string
    Transport *Transport

    // eventTarget holds an EventTarget
    eventTarget *EventTarget

    HandleFn func(EventReturnType) (error, EventReturnType)
}

func (c *DefaultEventController) GetEventTarget() *EventTarget {
    if c.eventTarget == nil {
		c.eventTarget = &EventTarget{
			Resource: c.Resource,
			Action: c.Action,
			Transport: c.Transport,
        }
    }

    return c.eventTarget
}

func (c *DefaultEventController) Handle(data EventReturnType) (error, EventReturnType) {
    return c.HandleFn(data)
}
