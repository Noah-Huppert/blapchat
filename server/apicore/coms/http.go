package coms

// HTTPCommunicator implements the Communicator interface for the HTTP
// transport method.
//
// This will allows clients to interact with the API via standard HTTP requests.
//
// Handlers will be registered under URLs with the following template:
//
//	/{{ .RootURL }}/{{ .Target.Version }}/{{ .Target.Resource }}
//
// This can be changed by providing a HndlrURLTmpl other than the
// DefaultHndlrURLTmpl.
//
// Handlers are registered with the HTTP request method that maps to the
// .Handler.Target.Action, based on the following mapping:
//
// 	- Create => POST
//	- Get    => GET
// 	- Update => UPDATE
//	- Delete => DELETE
//
// This can be changed by providing a HndlrActionMap other than the
// DefaultHndlrActionMap
type HTTPCommunicator struct {
	// HndlrURLTmpl is the Go template which will be evaluated to determine
	// which URL a handler should be registered under.
	//
	// The template will have the following data variables available:
	//
	//	- All exported HTTPCommunicator fields with no prefix (ex.,
	// 	  HTTPCommunicator.RootURL can be accessed at {{ .RootURL }}
	// 	- The Handler's Target which is being evaluated is
	//	  available under the `.Target` variable
	HndlrURLTmpl string

	// HndlrActionMap is a map with keys which are Target.Action values,
	// and HTTP verb values.
	//
	// This map will be used to determine which HTTP verb a handler will be
	// registered with for its route.
	HndlrActionMap map[string]string

	// RootURL is the URL which all handlers should be registered under
	RootURL string

	// server is the http.Server used to receive HTTP requests. It will be
	// set to nil if no http.Server has been created
	server *http.Server
}

// HTTPVerbPost holds the value which will be used to represent the Post HTTP
// verb
const HTTPVerbPost string = "POST"

// HTTPVerbGet holds the value which will be used to represent the Get HTTP verb
const HTTPVerbGet string = "GET"

// HTTPVerbUpdate holds the value which will be used to represent the Update
// HTTP verb
const HTTPVerbUpdate string = "UPDATE"

// HTTPVerbDelete holds the value which will be used to represent the Delete
// HTTP verb
const HTTPVerbDelete string = "DELETE"

// DefaultHndlrURLTmpl holds the default HndlrURLTmpl value
const DefaultHndlrURLTmpl string = "/{{ .RootURL }}/{{ .Target.Version }}/{{ .Target.Resource }}"

// DefaultHndlrActionMap holds the default HndlrActionMap value
const DefaultHndlrActionMap map[string]string = map[string]string{
	hdnlrs.actions.CreateAction: HTTPVerbPost,
	hndlrs.actions.GetAction:    HTTPVerbGet,
	hndlrs.actions.UpdateAction: HTTPVerbUpdate,
	hndlrs.actions.DeleteAction: HTTPVerbDelete,
}

// Communicate implements Communicator.Communicate for HTTPCommunicator.
//
// It will start a http.Server which will receive requests
func (c HTTPCommunicator) Communicate(invoker Invoker) <-chan error {
	errs := make(chan error)

	// Check a HTTP server hasn't been created yet
	if c.server != nil {
		errs <- errors.New("a server already exists")
		return errChan
	}

	// Create http.Server
	server := http.Server{handler: c}

	// Start http.Server
	if err := server.ListenAndServe(); err != nil {
		errs <- fmt.Errof("error starting server: %s", err.Error())
		return errChan
	}
}

// Stop implements Communicator.Stop for HTTPCommunicator.
//
// It will stop the stored c.server within the given Context
func (c HTTPCommunicator) Stop(ctx context.Context) <-chan error {

}
