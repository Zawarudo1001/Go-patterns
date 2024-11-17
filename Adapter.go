// Паттерн Adapter

package main

import "fmt"

// Интерфейс Target
type Target interface {
	Request() string
}

// Структура Adaptee с несовместимым методом
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request from Adaptee"
}

// Структура Adapter, которая реализует интерфейс Target
type Adapter struct {
	adaptee *Adaptee
}

// Реализация метода Request в адаптере
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	// Создаем экземпляр Adaptee
	adaptee := &Adaptee{}

	// Создаем адаптер, который использует Adaptee
	adapter := &Adapter{adaptee: adaptee}

	// Клиент использует адаптер
	fmt.Println(adapter.Request()) // Вывод: Specific request from Adaptee
}
