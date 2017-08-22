package util

var actionToVerbs map[string]string = map[string]string{
    "get": "GET",
    "create": "POST",
    "update": "PUT",
    "delete": "DELETE",
}

type HTTPLoader struct {}

func (l HTTPLoader) GetTransport() *Transport {
    return TransportHTTP
}

func (l HTTPLoader) Load(h *EventHandler) error {

}
