package goflux

func Register(identity interface{}, listener chan *Action) {
	Dispatcher.register <- newFlux(identity, listener)
}

func UnRegister(identity interface{}, listener chan *Action) {
	Dispatcher.unRegister <- newFlux(identity, listener)
}

func Send(name, from, to, payload interface{}) {
	Dispatcher.sendAction <- newAction(name, from, to, payload)
}
