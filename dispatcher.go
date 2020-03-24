// Copyright (C) 2019 JohnnyChu <chuhsun@gmail.com>
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
    "sync"
)

const (
    ActionFluxRegistered = "ActionFluxRegistered"
)

var Dispatcher = &dispatcher{
    register:   make(chan *flux, 1024),
    unRegister: make(chan *flux, 1024),
    async:      make(chan *Action, 1024),
    sync:       make(chan *Action, 1024),
    rwMutex:    &sync.RWMutex{},
}

type dispatcher struct {
    register   chan *flux
    unRegister chan *flux
    async      chan *Action
    sync       chan *Action
    rwMutex    *sync.RWMutex
}

func (d *dispatcher) start(workers int) {
    listenerGroup := make(map[interface{}]map[chan *Action]interface{})

    if workers <= 0 {
        workers = 10
    }

    for i := 0; i < workers; i++ {
        go Dispatcher.run(listenerGroup)
    }
}

func (d *dispatcher) run(listenerGroup map[interface{}]map[chan *Action]interface{}) {
    for {
        select {
        case flux := <-d.register:
            d.rwMutex.Lock()

            if listeners, ok := listenerGroup[flux.identity]; ok {
                listeners[flux.listener] = flux.identity
            } else {
                listeners := make(map[chan *Action]interface{})
                listeners[flux.listener] = flux.identity
                listenerGroup[flux.identity] = listeners
            }

            flux.listener <- newAction(ActionFluxRegistered, "", flux.listener, "")

            d.rwMutex.Unlock()

        case flux := <-d.unRegister:
            d.rwMutex.Lock()

            if listeners, ok := listenerGroup[flux.identity]; ok {
                delete(listeners, flux.listener)
                close(flux.listener)
                flux.listener = nil
            } else {
                log.Println("UnRegister: Can't find the flux listener.")
            }

            d.rwMutex.Unlock()

        case action := <-d.async:
            d.rwMutex.RLock()

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

            d.rwMutex.RUnlock()

        case action := <-d.sync:
            d.rwMutex.RLock()

            if listeners, ok := listenerGroup[action.to]; ok {
                for listener := range listeners {
                    listener <- action

                    if len(listener) >= cap(listener) {
                        log.Println("SendAction: The flux listener is overflow.", action)
                    }
                }
            } else {
                log.Println("SendAction: Can't find the flux listener.")
                action.sync <- false
            }

            d.rwMutex.RUnlock()
        }
    }
}
