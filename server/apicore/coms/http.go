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
// This can be changed by providing a HndlrUrlTmpl other than the
// DefaultHndlrUrlTmpl.
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
	// HndlrUrlTmpl is the Go template which will be evaluated to determine
	// which URL a handler should be registered under.
	//
	// The template will have the following data variables available:
	//
	//	- All exported HTTPCommunicator fields with no prefix (ex.,
	// 	  HTTPCommunicator.RootURL can be accessed at {{ .RootURL }}
	// 	- The Handler's Target which is being evaluated is
	//	  available under the `.Target` variable
	HndlrUrlTmpl string

	// HndlrActionMap TODO

	// RootURL is the URL which all handlers should be registered under
	RootURL string
}

// Communicate implements Communicator.Communicate for HTTPCommunicator.
//
// It will start a http.Server
