package controllers

import "github.com/Noah-Huppert/blapchat/server/handlers"

// Controller interface provides information about a group of handlers
type Controller interface {
    // Handlers returns all the handlers a controller groups
    Handlers() []handlers.Handler
}
