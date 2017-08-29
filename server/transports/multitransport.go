package transports

// MultiTransport routes register calls to transports in the .transports map based on HandlerTarget.TransportRestrictions
type MultiTransport struct {
    transports map[TransportType]*Transport
}


