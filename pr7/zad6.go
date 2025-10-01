package main

import (
	"fmt"
	"sync"
)

type EventBus struct {
	subscribers map[string][]func(interface{})
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(interface{})),
	}
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	eb.subscribers[event] = append(eb.subscribers[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	
	if handlers, exists := eb.subscribers[event]; exists {
		for _, handler := range handlers {
			handler(data)
		}
	}
}

func main() {
	podpiska := NewEventBus()
	
	podpiska.Subscribe("message", func(data interface{}) {
		fmt.Printf("Подписчик 1 получил: %v\n", data)
	})
	
	podpiska.Subscribe("message", func(data interface{}) {
		fmt.Printf("Подписчик 2 получил: %v\n", data)
	})

	podpiska.Subscribe("error", func(data interface{}) {
		fmt.Printf("Обработчик ошибок получил: %v\n", data)
	})

	podpiska.Publish("message", "Учить go реально интересно)")
	podpiska.Publish("error", "Произошла ошибка")
	podpiska.Publish("unknown", "Это событие никому не достанется")
}