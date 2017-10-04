package evnt

// Event holds information ahout a request received from a client. The concept
// of an event is used to abstract away the actual communication method of the
// client request and response.
//
// See wiki/technical/events/Events.md for more information.
type Event struct {
	// Target holds information used to determine which action to invoke
	Target Target

	// Payload holds any data which was provided with event when it was
	// received.
	Payload interface{}
}

// Merge returns a new Event which is a combination of a & b:
//
// 	a) the Event the method was called on
//	b) the Event provided as an argument to the method
func (e Event) Merge(o Event) Event {

}
