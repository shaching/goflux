# goflux

## What is FLUX? See Here https://facebook.github.io/flux/

# Step 1. Install goflux.

````
go get github.com/shaching/goflux
``````

# Step 2. Register your go channels.
* The ***listener*** is must be ***chan \*Action*** type in your type struct.
* Notice: If you use block channel, you must implement ***Step 4.***
    
```golang
type foo struct {
    listener  chan *goflux.Action
}

f := &foo{
    listener:  make(chan *goflux.Action),
}

goflux.Register(f.listener)
```

```golang
type bar struct {
    listener  chan *goflux.Action
}

b := &bar{
    listener:  make(chan *goflux.Action),
}

goflux.Register(b.listener)
```

# Step 3. Create a action.
* Action name ***foo*** is ***interface{}*** type.
* Action payload ***bar*** is ***interface{}*** type.

```golang
goflux.Send(goflux.NewAction(foo1, bar1))
or
goflux.Send(goflux.NewAction(foo2, bar2))
```

# Step 4. Implement your flux store with goroutine listener func.
* Notice: You can filter action here.

```golang
go f.store()

func (f *foo) store() {
    for {
        action, ok := <-f.listener
	if !ok {
            return
        }
       // will receive foo1, foo2 action and bar1, bar2 payload
       // you can filter received action here
    }
}
```

```golang
go b.store()

func (b *bar) store() {
    for {
        action, ok := <-b.listener
        if !ok {
            return
        }
        // will receive foo1, foo2 action and bar1, bar2 payload
        // you can filter received action here
    }
}
```

# Step 5. Also, unregister when you wan't use.
* Notice: DON'T manual close channel by yourself, this func will automatic close it. 

```golang
goflux.UnRegister(f.listener)
or
goflux.UnRegister(b.listener)
```
