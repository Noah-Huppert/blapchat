package util

import (
    "github.com/labstack/echo"
    "fmt"
)

type EchoFn func(string, echo.HandlerFunc, ...echo.MiddlewareFunc)

// verbInfo is a struct which holds metadata information about Http verbs and their programatic usage in the http loader
type verbInfo struct {
    // HttpVerb is the "official" name of the http verb. ex., GET, DELETE, POST
    HttpVerb string

    // PathAddition is an optional string that will be added to the end of the path which the handler will be registered at.
    // This is used for http verbs where there is no direct mapping and a sub-url is needed (such as /users/invite,
    // "/invite" is the path addition).
    //
    // If provided it must start with a slash ("/"). If not needed set to an empty string
    PathAddition string

    // EchoRegisterFn is a function which must return the the echo http framework function for that particular http verb
    // when called
    EchoRegisterFn func (e echo.Echo) EchoFn
}

// actionToVerbs is a map. String keys are event action values. Values are verbInfo structs
var actionToVerbs map[string]verbInfo = map[string]verbInfo{
    // Standard verbs
    "get": verbInfo{"GET", "", func(e echo.Echo) EchoFn {
        return e.GET
    }},
    "create": verbInfo{"POST", "", func(e echo.Echo) EchoFn {
        return e.POST
    }},
    "update": verbInfo{"PUT", "", func(e echo.Echo) EchoFn {
        return e.PUT
    }},
    "delete": verbInfo{"DELETE", "", func(e echo.Echo) EchoFn {
        return e.DELETE
    }},

    // Special actions
    "list": verbInfo{"GET", "/list", func(e echo.Echo) EchoFn {
        return e.GET
    }},
    "view": verbInfo{"POST", "/view", func(e echo.Echo) EchoFn {
        return e.POST
    }},
}

type HTTPLoader struct {
    echo *echo.Echo
}

func (l HTTPLoader) GetTransport() *Transport {
    return TransportHTTP
}

func (l HTTPLoader) Load(resource string, h *EventHandler) error {
    // Check action is known
    if _, ok := actionToVerbs[h.GetAction()]; !ok {
        // If unknown, error
        return fmt.Errorf("Unknown action: \"%s\"", h.GetAction())
    }

    // Save to Handlers registry
    AddHandler(resource, h)


    // Determine path
    verb := actionToVerbs[h.GetAction()]
    pathMod := ""

    if len(verb) > 1 {
        pathMod = verb[1]
    }

    path := fmt.Sprintf("/%s/%s%s", resource, h.GetAction(), pathMod)

    // Register with echo
    echo.GET
}
