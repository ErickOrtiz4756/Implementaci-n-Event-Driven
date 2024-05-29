package govent

// Event puede tomar todos los argumentos de un evento
type Event interface{}

// EventType es el tipo de evento para un evento
type EventType int

// EventObject contiene el evento y el tipo de evento
type EventObject struct {
	EventType
	Event
}

// EventHandler es el controlador de eventos y toma cualquier argumento
type EventHandler func(args Event)

// Registrar tipos de eventos, esto es solo un formato
const (
	NoobEvent EventType = iota
)
