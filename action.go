package goflux

type Action struct {
	name    interface{}
	payload interface{}
}

func NewAction(name, payload interface{}) *Action {
	return &Action{
		name:    name,
		payload: payload,
	}
}

func (a *Action) Name() interface{} {
	return a.name
}

func (a *Action) Payload() interface{} {
	return a.payload
}
