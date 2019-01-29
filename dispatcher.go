package goflux

var Dispatcher = &dispatcher{
	register:   make(chan chan *Action),
	unRegister: make(chan chan *Action),
	sendAction: make(chan *Action, 10),
}

type dispatcher struct {
	register   chan chan *Action
	unRegister chan chan *Action
	sendAction chan *Action
}

func init() {
	go Dispatcher.run()
}

func (d *dispatcher) run() {
	listeners := make(map[chan *Action]chan *Action)

	for {
		select {
		case listener := <-d.register:
			listeners[listener] = listener
		case listener := <-d.unRegister:
			delete(listeners, listener)
			close(listener)
			listener = nil
		case action := <-d.sendAction:
			for listener := range listeners {
				listener <- action
			}
		}
	}
}
