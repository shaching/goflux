# goflux
> Callback is usually used between different objects' data transmission and it will cause callback hell. So, decoupling is must necessary and goflux can solve the problem.

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
    address   string
    listener  chan *goflux.Action
}

e := &earth{
    address:   "Earth",
    listener:  make(chan *goflux.Action),
}
```

```golang
type mars struct {
    address   string
    listener  chan *goflux.Action
}

m := &mars{
    address:   "Mars",
    listener:  make(chan *goflux.Action),
}
```

# Step 3. Create a action including below four parameters.
* name: action name.
* from: where action comes from.
* to: where action goes to.
* payload: data.

```golang
goflux.Send("ActionCall", "Earth", "Mars", "Hi", "Mom")
or
goflux.Send("ActionReCall", "Mars", "Earth", "Hello", "Dad")
```

# Step 4. Implement your flux store with goroutine listener func.
* Notice: Register listener inner goroutine func.
* Notice: You can filter action here.

```golang
go e.store() <-- Start goroutine.

func (e *earth) store() {
    goflux.Register(e.address, e.listener) <-- Register Flux.

    for {
        action, ok := <-e.listener
        
        if !ok {
            return
        }
        
        // action.Name().(string)        --> ActionReCall
        // action.From().(string)        --> Mars
        // action.To().(string)          --> Earth
        // action.Payload()[0].(string)  --> Hello
        // action.Payload()[1].(string)  --> Dad

        // you can handle received action here
    }
}
```

```golang
go m.store() <-- Start goroutine.

func (m *mars) store() {
    goflux.Register(m.address, m.listener) <-- Register Flux.
    
    for {
        action, ok := <-m.listener
        
        if !ok {
            return
        }
        
        // action.Name().(string)        --> ActionCall
        // action.From().(string)        --> Earth
        // action.To().(string)          --> Mars
        // action.Payload()[0].(string)  --> Hi
        // action.Payload()[1].(string)  --> Mom
        
        // you can handle received action here
    }
}
```

# Step 5. Also, unregister when you wan't use.
* Notice: DON'T manual close channel by yourself, goflux will automatic close it. 

```golang
goflux.UnRegister(e.address, e.listener)
or
goflux.UnRegister(m.address, m.listener)
```
