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
	var actionLoaderErrs []error = []error{}

	// Creator
	if l.Creator != nil {
		target := l.ParentTarget.Merge(l.CreatorTarget)
		loader := CreatorLoader{l.Creator, target}

		actionLoaderErrs = l.LoadOtherLoader(actionLoaderErrs, "Creator", loader, r)
	}

	// Getter
	if l.Getter != nil {
		target := l.ParentTarget.merge(l.GetterTarget)
		loader := GetterLoader{l.Getter, target}

		actionLoaderErrs = l.LoadOtherLoader(actionLoaderErrs, "Getter", loader, r)
	}

	// Updater
	if l.Updater != nil {
		target := l.ParentTarget.Merge(l.UpdaterTarget)
		loader := UpdaterLoader{l.Updater, target}

		actionLoaderErrs = l.LoadOtherLoader(actionLoaderErrs, "Updater", loader, r)
	}

	// Deleter
	if l.Delter != nil {
		target := l.ParentTarget.Merge(l.DeleterTarget)
		loader = DeleterLoader{l.Deleter, target}

		actionLoaderErrs = l.LoaderOtherLoader(actionLoaderErrs, "Deleter", loader, r)
	}

	// Check if any load errors
	if len(actionLoaderErrs) > 0 {
		// If errors, return
		return fmt.Errorf("error loading actions: %s", actionLoaderErrrs)
	}

	// If no errors, success
	return nil
}

// LoadOtherLoader calls a provided Loaders Loader.load method and the provided
// `r` Registry.
//
// If an error occurs, it is appended to the `errs` array. Using the `name`
// argument to identify which Loader caused the error.
//
// A new `errs` array is returned by this method, with any potential errors
// that occurred while calling the loader.
func (l Loader) LoadOtherLoader(errs []error, name string, other Loader, r Registry) []error {
	if err := other.Load(r); err != nil {
		errs = append(errs, fmt.Errorf("error loading %s: %s", name, err.Error()))
	}

	return errs
}
