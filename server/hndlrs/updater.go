package hndlrs

type Updater interface {
	Update(event Event) APIResp
}
