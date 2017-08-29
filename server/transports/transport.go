package transports

// TransportType is a type alias for the name of a transport
type TransportType string
const (
    Http TransportType = "http"
    Fcm TransportType = "fcm"
)

// Transport defines methods for managing handler registration for a transport. This allows the backend of a transport
// to be configured in the defined registers.
//
// Transport is a term for a system which passes information. The Transport interface defines the methods necessary to
// register handlers with a system which passes information.
//
// In the register handler method one should use a backend library apis to register the appropriate handlers.
//
// Ex:
//     - Transport system named Foo
//         - Foo is an event system
//         - Foo clients can subscribe to message topics
//     - Golang library named `foot` provides methods for using Foo
//         - `foot.NewClient(authToken string) (error, foo.Client)` to create a foot client
//         - `foo.Client.Sub(topic string, handler func(message string) (err error, response string)) error` to
//            subscribe to and handle messages
//
//     const Foo TransportType = "foo"
//
//     type FooTransport struct {
//         fClient foot.Client
//     }
//
//     func (t FooTransport) Name() TransportType {
//         return Foo
//     }
//
//     func (t FooTransport) RegisterHandler(h Handler) error {
//         // Use foo client to subscribe to topic and simply echo all messages back to sender
//         return t.fClient.Sub(strings.Join(h.Path, "."), func(message string) (err error, response string) {
//             response := message
//             return
//         })
//     }

type Transport interface {
    // Name returns the TransportType which the transport operates for
    Name() TransportType

    // RegisterHandler will add the handler to the transport's proper backend systems to allow requests to be received
    // by the handler.
    //
    // RegisterHandler must return with an error if the provided handler has a HandlerTarget with a .TransportRestrictions
    // value that prohibits the transports type. This can be checked with EventTarget.IncludesTransport(<Transport Type>)
    RegisterHandler(h Handler) error
}
