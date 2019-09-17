// Copyright (C) 2019 JohnnyChu(chuhsun@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goflux

import (
	"log"
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
	listenerGroup := make(map[interface{}]map[chan *Action]interface{})

	for {
		select {
		case flux := <-d.register:
			if listeners, ok := listenerGroup[flux.identity]; ok {
				listeners[flux.listener] = flux.identity
			} else {
				listeners := make(map[chan *Action]interface{})
				listeners[flux.listener] = flux.identity
				listenerGroup[flux.identity] = listeners
			}
		case flux := <-d.unRegister:
			if listeners, ok := listenerGroup[flux.identity]; ok {
				delete(listeners, flux.listener)
				close(flux.listener)
				flux.listener = nil
			} else {
				log.Println("UnRegister: Can't find the flux listener.")
			}
		case action := <-d.sendAction:
			if listeners, ok := listenerGroup[action.to]; ok {
				for listener := range listeners {
					listener <- action

					if len(listener) >= cap(listener) {
						log.Println("SendAction: The flux listener is overflow.", action)
					}
				}
			} else {
				log.Println("SendAction: Can't find the flux listener.")
			}
		}
	}
}
