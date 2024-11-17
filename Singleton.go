//Паттерн Singleton

package main

import (
	"fmt"
	"sync"
)

// Структура синглтона
type Singleton struct{}

// Переменная для хранения единственного экземпляра
var instance *Singleton
var once sync.Once

// Приватный конструктор
func newSingleton() *Singleton {
	return &Singleton{}
}

// Публичный метод для получения единственного экземпляра
func GetInstance() *Singleton {
	once.Do(func() {
		instance = newSingleton()
	})
	return instance
}

// Метод для демонстрации работы синглтона
func (s *Singleton) DoSomething() {
	fmt.Println("Doing something in the singleton instance.")
}

// Клиентский код
func main() {
	singleton1 := GetInstance()
	singleton1.DoSomething()

	singleton2 := GetInstance()
	singleton2.DoSomething()

	// Проверка, что оба экземпляра равны
	fmt.Println(singleton1 == singleton2) // Вывод: true
}
