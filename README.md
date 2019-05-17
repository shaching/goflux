# goflux

## What is FLUX? See Here https://facebook.github.io/flux/

# Step 1. Install goflux.

````
go get github.com/shaching/goflux
``````

# Step 2. Construct go channels.
* The ***listener*** is must be ***chan \*Action*** type in your type struct.
* Notice: Remember implement ***Step 4.***
    
```golang
type earth struct {
    listener  chan *goflux.Action
}

&earth{
    listener:  make(chan *goflux.Action),
}
```

```golang
type mars struct {
    listener  chan *goflux.Action
}

&mars{
    listener:  make(chan *goflux.Action),
}
```

# Step 3. Create a action including below four parameters.
* name: action name.
* from: where action comes from.
* to: where action goes to.
* payload: data.

```golang
goflux.Send("ActionCall", "Earth", "Mars", "Hi, Mars.")
or
goflux.Send("ActionReCall", "Mars", "Earth", "Hello, Earth.")
```

# Step 4. Implement your flux store with goroutine listener func.
* Notice: Register listener inner goroutine func.
* Notice: You can filter action here.

```golang
go f.store()

func (e *earth) store() {
    goflux.Register(e.listener)

    for {
        action, ok := <-e.listener
        
        if !ok {
            return
        }
        // will receive action "ActionReCall" and "Hello, Earth." payload
        // you can handle received action here
    }
}
```

```golang
go b.store()

func (m *mars) store() {
    goflux.Register(m.listener)
    
    for {
        action, ok := <-m.listener
        
        if !ok {
            return
        }
        // will receive action "ActionCall" and "Hi, Mars." payload
        // you can handle received action here
    }
}
```

# Step 5. Also, unregister when you wan't use.
* Notice: DON'T manual close channel by yourself, goflux will automatic close it. 

```golang
goflux.UnRegister(e.listener)
or
goflux.UnRegister(m.listener)
```
