package evnt

// Target provides information used to determine which action to invoke when an
// event is received. This information includes:
//
//     - Target resource
//     - Target action
//     - Target API version
//
// See the wiki/technical/events/Events.md#Target page for more information
type Target interface {
	// Resource returns the name of the resource targeted to run an action
	Resource() string

	// Action returns the name of the action targeted to run
	Action() string

	// Version returns the SemVer v2 version of the API which is targeted
	Version() SemVer
}
