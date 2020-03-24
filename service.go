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

func Init(workers int) {
    Dispatcher.start(workers)
}

func Register(identity interface{}, listener chan *Action) {
    Dispatcher.register <- newFlux(identity, listener)
}

func UnRegister(identity interface{}, listener chan *Action) {
    Dispatcher.unRegister <- newFlux(identity, listener)
}

func Async(name, from, to interface{}, payload ...interface{}) {
    Dispatcher.async <- newAction(name, from, to, payload)
}

func Sync(name, from, to interface{}, payload ...interface{}) chan bool {
    action := newAction(name, from, to, payload)
    Dispatcher.sync <- action
    return action.sync
}
