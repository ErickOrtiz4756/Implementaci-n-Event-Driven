Arquitectura Event driven en Go utilizando canales

Registra nuevos tipos de eventos - 
```go
const (
	MyMessage govent.EventType = iota
)
```

Realiza nuevos eventos -
```go
type MessageEvent struct {
	Message string
}
```

Hace una funcion de consumidor para el evento -
```go
func ShowMessage(e govent.Event) {
	if e, ok := e.(MessageEvent); ok {
		fmt.Println(e.Message)
	}
}
```

Iniciar un nuevo editor de eventos -
```go
publisher := govent.NewEventPublisher()
```

Suscribir eventos al editor de eventos -
```go
govent.Subscribe(ShowMessage, MyMessage)
```

Publica eventos - 
```go
publisher <- govent.EventObject{
		EventType: MyMessage,
		Event:     MessageEvent{Message: "Hello World"},
	}
```

## Run
```
cd example
go run *.go
```
