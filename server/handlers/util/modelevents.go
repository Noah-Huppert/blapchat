package util

// ModelEvents is a struct which holds all the event handlers a model has
type ModelEvents struct {
    Handlers []EventHandler
}

// NewModelEvents returns a ModelEvents struct with the provided EventHandler array
func NewModelEvents(handlers []EventHandler) ModelEvents {
    return ModelEvents{handlers}
}
