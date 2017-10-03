package hndlrs

type Creator interface {
	Create(event Event) APIResp
}
