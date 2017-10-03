package hndlrs

type Getter interface {
	Get(event Event) APIResp
}
