package actions

// Loader loads the provided: Creator, Getter, Updater, and / or Deleter
// if they are not nil. By using the respective: CreatorLoader, GetterLoader,
// UpdaterLoader, and / or DeleterLoader interface(s).
type Loader struct {
	// ParentTarget holds the default Target value which will be used
	// to register the provided action interfaces. The ParentTarget.Action
	// field will not be evuated, as it changes for each action.
	ParentTarget Target

	// Creator holds a Creator interface if one should be loaded. Otherwise
	// value should not be provided
	Creator Creator

	// CreatorTarget will override ParentTarget values with non-empty values
	// of its own
	CreatorTarget Target

	// Getter holds a Getter interface if one should be loaded. Otherwise
	// value should not be provided
	Getter Getter

	// GetterTarget will override ParentTarget values with non-empty values
	// of its own
	GetterTarget Target

	// Updater holds a Updater interface if one should be loaded. Otherwise
	// value should not be provided
	Updater Updater

	// UpdaterTarget will override ParentTarget values with non-empty values
	// of its own
	UpdaterTarget Target

	// Deleter holds a Deleter interface if one should be loaded. Otherwise
	// value should not be provided
	Deleter Deleter

	// DeleterTarget will override ParentTarget values with non-empty values
	// of its own
	DeleterTarget Target
}

// Load implements the hndlrs.Loader.Load function for Loader
func (l Loader) Load(r Registry) error {

}
