package goflux

type flux struct {
	identity interface{}
	listener chan *Action
}

func newFlux(identity interface{}, listener chan *Action) *flux {
	return &flux{
		identity: identity,
		listener: listener,
	}
}
