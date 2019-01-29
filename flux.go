package goflux

func Register(listener chan *Action) {
	Dispatcher.register <- listener
}

func UnRegister(listener chan *Action) {
	Dispatcher.unRegister <- listener
}

func Send(action *Action) {
	Dispatcher.sendAction <- action
}
