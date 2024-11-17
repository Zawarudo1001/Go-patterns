//Паттерн Strategy

package main

import "fmt"

// Интерфейс стратегии
type Strategy interface {
	Execute(a, b int) int
}

// Конкретная стратегия для сложения
type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

// Конкретная стратегия для вычитания
type SubtractStrategy struct{}

func (s *SubtractStrategy) Execute(a, b int) int {
	return a - b
}

// Конкретная стратегия для умножения
type MultiplyStrategy struct{}

func (s *MultiplyStrategy) Execute(a, b int) int {
	return a * b
}

// Контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

// Метод для установки стратегии
func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

// Метод для выполнения стратегии
func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

// Клиентский код
func main() {
	context := &Context{}

	// Использование стратегии сложения
	context.SetStrategy(&AddStrategy{})
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	// Использование стратегии вычитания
	context.SetStrategy(&SubtractStrategy{})
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))

	// Использование стратегии умножения
	context.SetStrategy(&MultiplyStrategy{})
	fmt.Println("10 * 5 =", context.ExecuteStrategy(10, 5))
}
