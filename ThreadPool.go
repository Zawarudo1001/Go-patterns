/*
Паттерн ThreadPool
*/

package main

import (
	"fmt"
	"sync"
)

// Определяем структуру задачи
type Task func()

// Определяем структуру Worker
type Worker struct {
	id         int
	taskQueue  chan Task
	wg         *sync.WaitGroup
}

// Определяем структуру ThreadPool
type ThreadPool struct {
	workers    []*Worker
	taskQueue  chan Task
	wg         sync.WaitGroup
}

// Создаем новый Worker
func NewWorker(id int, taskQueue chan Task, wg *sync.WaitGroup) *Worker {
	return &Worker{id: id, taskQueue: taskQueue, wg: wg}
}

// Запускаем Worker
func (w *Worker) Start() {
	go func() {
		for task := range w.taskQueue {
			task() // Выполняем задачу
			w.wg.Done() // Уменьшаем счетчик WaitGroup
		}
	}()
}

// Создаем новый ThreadPool
func NewThreadPool(numWorkers int) *ThreadPool {
	taskQueue := make(chan Task)
	pool := &ThreadPool{taskQueue: taskQueue}

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i, taskQueue, &pool.wg)
		worker.Start()
		pool.workers = append(pool.workers, worker)
	}

	return pool
}

// Добавляем задачу в ThreadPool
func (p *ThreadPool) AddTask(task Task) {
	p.wg.Add(1) // Увеличиваем счетчик WaitGroup
	p.taskQueue <- task
}

// Завершаем работу ThreadPool
func (p *ThreadPool) Shutdown() {
	close(p.taskQueue) // Закрываем очередь задач
	p.wg.Wait() // Ждем завершения всех задач
}

func main() {
	pool := NewThreadPool(3) // Создаем ThreadPool с 3 рабочими потоками

	// Добавляем задачи в ThreadPool
	for i := 0; i < 10; i++ {
		i := i // Создаем локальную переменную для захвата
		pool.AddTask(func() {
			fmt.Printf("Task %d is running\n", i)
		})
	}

	pool.Shutdown() // Завершаем работу ThreadPool
}
