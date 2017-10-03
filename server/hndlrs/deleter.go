package hndlrs

type Deleter interface {
	Delete(event Event) APIResp
}
