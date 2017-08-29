package handlers

import "github.com/Noah-Huppert/blapchat/server/transports"

type TestHandler struct {}

// NewTestHandler creates a new TestHandler instance. It returns type *Handler
func NewTestHandler() transports.Handler {
    return new(TestHandler)
}

func (h TestHandler) Target() transports.HandlerTarget {
    return transports.StandardHandlerTarget(
        []string{"namespace"},
        transports.Get,
    )
}

func (h TestHandler) Handle(data transports.HandlerData) (error, transports.HandlerData) {
    var resp transports.HandlerData
    resp["data"] = data

    return nil, resp
}
