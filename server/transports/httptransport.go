package transports

import (
    "net/http"
    "fmt"
    "strings"
)

// HttpTransport implements the Transport interface for Http requests
type HttpTransport struct {
    // mux is the ttp.ServeMux to register routes with
    mux *http.ServeMux

    pathHandlers map[string]*HttpPathHandler
}

// NewHttpTransport creates a new instance of a HttpTransport with the provided ServeMux. As well as an empty
// .pathHandlers map
func NewHttpTransport(mux *http.ServeMux) *HttpTransport {
    return &HttpTransport{mux, map[string]*HttpPathHandler{}}
}

// Transport interface methods
func (t *HttpTransport) Name() TransportType {
    return Http
}

// RegisterHandler
func (t *HttpTransport) RegisterHandler(h Handler) error {
    // Check if transport included in handler target
    target := h.Target()
    targetStr := target.String()

    if !target.IncludesTransport(t.Name()) {
        return fmt.Errorf("HandlerTarget does not include transport type \"%s\"", t.Name())
    }

    // Check HttpPathHandler for path exists
    path := PathToSlashed(target.Path)

    if pathHandler, exists := t.pathHandlers[path]; !exists {
        // If one doesn't exist, make one
        t.pathHandlers[path] = NewHttpPathHandler(path)

        // And register with ServeMux
        if err := pathHandler.MuxRegister(*t.mux); err != nil {
            return fmt.Errorf("Error registering HttpPathHandler at path \"%s\", error: \"%s\"", path, err.Error())
        }
    }

    // Check if handler is already registered
    handlers := t.pathHandlers[path].Handlers()

    if _, exists := handlers[targetStr]; exists {
        // If exists, return already registered error
        return fmt.Errorf("Handler already registered for EventTarget: %s", target)
    }

    // If doesn't exist, added to handlers map so HttpPathHandler will handle that action
    handlers[targetStr] = h
    return nil
}

// Helpers
// VerbForAction returns the Http veb for a provided action. Or an error if the provided action is not known.
func VerbForAction(action ActionType) (error, string) {
    switch action {
    case Create:
        return nil, "POST"
    case Get:
        return nil, "GET"
    case Update:
        return nil, "PUT"
    case Delete:
        return nil, "DELETE"
    default:
        return fmt.Errorf("Unknown action \"%s\"", action), ""
    }
}

// PathToSlashed returns the concatenated version of the provided path, separated by slashes: "/".
func PathToSlashed(path []string) string {
    return strings.Join(path, "/")
}

// HttpPathHandler is a struct which implements the golang http handler interface. This handler will receive http
// requests and look at the .handlers array for Handlers to respond to request
type HttpPathHandler struct {
    // Path is the Http uri for which this handler should respond to requests for. Path is used to make sure application
    // handlers being added belong to this http path handler
    Path string

    // handlers holds the Handler interfaces which the path handler can use to respond to requests
    // handlers is a map with string keys. These string keys are string representations of HandleTarget structs. Values
    // are Handler pointers. The value is a pointer to the handler that should be used to respond to the request
    handlers map[string]Handler
}

// NewHttpPathHandler creates a new HttpPathHandler for the specified path, with an empty handlers map
func NewHttpPathHandler(path string) *HttpPathHandler {
    return &HttpPathHandler{path, map[string]Handler{}}
}

// Handlers returns the contents of the .handlers array
func (h HttpPathHandler) Handlers() map[string]Handler {
    return h.handlers
}

// MuxRegister adds this HttpPathHandler to the specified ServeMux. Returns nil if successful, error if not.
func (h HttpPathHandler) MuxRegister(mux http.ServeMux) (err error) {
    defer func() {
        // Recover if path is already registered with ServeMux
        if recover() != nil {
            err = fmt.Errorf("Path \"%s\" already registered on provided ServeMux", h.Path)
        }
    }()

    // Attempt to register handler at path, if it panics then the defer function above will recover and return an error
    mux.Handle(h.Path, h)
    return nil
}

// AddHandler adds the provided handler to the .handlers array. It will check that the provided handler has the correct
// path for this HttpPathHandler and error if it does not.
func (h *HttpPathHandler) AddHandler(handler Handler) error {
    target := handler.Target()
    targetStr := target.String()

    // Check if handler exists
    if _, exists := h.handlers[targetStr]; exists {
        return fmt.Errorf("Handler already exists")
    }

    // If doesn't exist, add
    h.handlers[targetStr] = handler
    return nil
}

func (h HttpPathHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL)
}
