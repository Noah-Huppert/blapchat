package evnt

// Target provides information used to determine which action to invoke when an
// event is received. This information includes:
//
//     - Target resource
//     - Target action
//     - Target API version
//
// See the wiki/technical/events/Events.md#Target page for more information
type Target struct {
	// Resource is the name of the resource targeted to run an action
	Resource string

	// Action is the name of the action targeted to run
	Action string

	// Version is the SemVer v2 version of the API which is targeted
	Version SemVer
}

// Merge returns a new Target which is a combination of a & b:
//
// 	a) the Target method was called on
//	b) the Target provided as an argument to the method
func (t Target) Merge(o Target) Target {
	// New target
	n := Target{t.Resource, t.Action, t.Version}

	// Other Target
	// Resource
	if o.Resource != "" {
		n.Resource = o.Resource
	}

	// Action
	if o.Action != "" {
		n.Action = o.Action
	}

	// Version
	// If other o.Version field is not 0, override
	if o.Version.Major != 0 {
		n.Version.Major = o.Version.Major
	}

	if o.Version.Minor != 0 {
		n.Version.Minor = o.Version.Minor
	}

	if o.Version.Patch != 0 {
		n.Version.Patch = o.Version.Patch
	}
}
