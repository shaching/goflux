package goflux

type Action struct {
	name    interface{}
	from    interface{}
	to      interface{}
	payload interface{}
}

func newAction(name, from, to, payload interface{}) *Action {
	return &Action{
		name:    name,
		from:    from,
		to:      to,
		payload: payload,
	}
}

func (a *Action) Name() interface{} {
	return a.name
}

func (a *Action) From() interface{} {
	return a.from
}

func (a *Action) To() interface{} {
	return a.to
}

func (a *Action) Payload() []interface{} {
	return a.payload.([]interface{})
}
