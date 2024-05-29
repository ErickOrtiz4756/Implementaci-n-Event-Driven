package main

import "github/ananto/govent"

// Mis tipos de eventos, registrándose en Govent
const (
	MyMessage govent.EventType = iota
	MyWeather govent.EventType = iota
)

// MessageEvent es un tipo de evento para mensajería
type MessageEvent struct {
	Message string
}

// WeatherEvent es un tipo de evento para condiciones meteorológicas
type WeatherEvent struct {
	Condition string
}
