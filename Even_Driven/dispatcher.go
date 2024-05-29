package govent

import "fmt"

var eventMap map[EventType][]EventHandler

func init() {
	eventMap = make(map[EventType][]EventHandler)
}

// Suscribirse a una función
func Subscribe(eH EventHandler, eT EventType) {
	if len(eventMap[eT]) == 0 {
		eventMap[eT] = make([]EventHandler, 0)
	}
	eventMap[eT] = append(eventMap[eT], eH)
}

// Publica el evento
func Publish(eT EventType, e Event) {
	for _, eH := range eventMap[eT] {
		eH(e)
	}
}

// Run es una goroutine para recibir y publicar eventos
func Run(publisher chan EventObject) {
	for {
		eventObject := <-publisher
		fmt.Println("Event received ", eventObject.EventType)
		Publish(eventObject.EventType, eventObject.Event)
	}
}

// NewEventPublisher crea un nuevo canal de editor para eventos
func NewEventPublisher() chan EventObject {
	publisher := make(chan EventObject)
	go Run(publisher)
	return publisher
}
