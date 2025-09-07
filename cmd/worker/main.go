package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

// WorkerStatus tracks what each worker is currently doing
type WorkerStatus struct {
	mu          sync.RWMutex
	workers     map[int]int  // worker ID -> current task ID
	lastUpdated map[int]bool // worker ID -> was just updated
}

func NewWorkerStatus() *WorkerStatus {
	return &WorkerStatus{
		workers:     make(map[int]int),
		lastUpdated: make(map[int]bool),
	}
}

func (ws *WorkerStatus) updateWorker(workerID, taskID int) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	// Reset all "just updated" flags
	for id := range ws.lastUpdated {
		ws.lastUpdated[id] = false
	}

	// Mark this worker as just updated
	ws.workers[workerID] = taskID
	ws.lastUpdated[workerID] = true
}

func (ws *WorkerStatus) printStatus(timestamp string) {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	// Clear screen before printing new status
	clearScreen()

	fmt.Printf("--------[%s]--------\n", timestamp)
	for i := 1; i <= 5; i++ {
		if taskID, exists := ws.workers[i]; exists {
			if ws.lastUpdated[i] {
				fmt.Printf("Worker %d started task %d (new)\n", i, taskID)
			} else {
				fmt.Printf("Worker %d started task %d\n", i, taskID)
			}
		}
	}
}

// clearScreen clears the console output
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Task represents a work item
type Task struct {
	ID    int
	Value int
}

// Result represents the output of a task
type Result struct {
	ID     int
	Value  int
	Result int
}

// Worker processes tasks from the task channel
func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup, status *WorkerStatus) {
	defer wg.Done()

	for task := range tasks {
		// Update worker status and print all workers
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		status.updateWorker(id, task.ID)
		status.printStatus(timestamp)

		// Random delay between 1-10 seconds to observe dynamic concurrency
		randomDelay := time.Duration(rand.Intn(10)+1) * time.Second
		time.Sleep(randomDelay)

		// Create a simple result (no square operation needed)
		result := Result{
			ID:     task.ID,
			Value:  task.Value,
			Result: 0, // No calculation needed
		}

		results <- result
	}
}

func main() {
	fmt.Println("Starting Worker Pool Demo")
	fmt.Println("=========================")

	const numTasks = 100
	const numWorkers = 5

	// Create channels
	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Create wait group for workers
	var wg sync.WaitGroup

	// Create worker status tracker
	status := NewWorkerStatus()

	// Start workers
	fmt.Printf("Starting %d workers...\n\n", numWorkers)

	// Start all workers simultaneously
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg, status)
	}

	// Give workers a moment to start, then begin sending tasks
	time.Sleep(100 * time.Millisecond)

	// Send tasks
	fmt.Printf("Sending %d tasks...\n\n", numTasks)
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- Task{
				ID:    i,
				Value: i,
			}
		}
		close(tasks)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Store results in a slice to maintain order
	resultSlice := make([]Result, numTasks)
	for result := range results {
		resultSlice[result.ID-1] = result
	}

	fmt.Println("\nWorker pool demo completed!")
}
