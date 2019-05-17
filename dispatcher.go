package goflux

import (
	"log"
	"sync"
)

var Dispatcher = &dispatcher{
	register:   make(chan *flux, 512),
	unRegister: make(chan *flux, 512),
	sendAction: make(chan *Action, 512),
}

type dispatcher struct {
	register   chan *flux
	unRegister chan *flux
	sendAction chan *Action
}

func init() {
	go Dispatcher.run()
}

func (d *dispatcher) run() {
	mutex := sync.Mutex{}
	listenerGroup := make(map[interface{}]map[chan *Action]interface{})

	for {
		mutex.Lock()
		select {
		case flux := <-d.register:
			if _, ok := listenerGroup[flux.identity]; ok {
				listeners := listenerGroup[flux.identity]
				listeners[flux.listener] = flux.identity
			} else {
				listeners := make(map[chan *Action]interface{})
				listeners[flux.listener] = flux.identity
				listenerGroup[flux.identity] = listeners
			}
		case flux := <-d.unRegister:
			if _, ok := listenerGroup[flux.identity]; ok {
				listeners := listenerGroup[flux.identity]
				delete(listeners, flux.listener)
				close(flux.listener)
				flux.listener = nil
			} else {
				log.Println("UnRegister: Can't find the flux listener.")
			}
		case action := <-d.sendAction:
			if _, ok := listenerGroup[action.to]; ok {
				for listener := range listenerGroup[action.to] {
					listener <- action

					if len(listener) >= cap(listener) {
						log.Println("SendAction: The flux listener is overflow.", action)
					}
				}
			} else {
				log.Println("SendAction: Can't find the flux listener.")
			}
		}
		mutex.Unlock()
	}
}
