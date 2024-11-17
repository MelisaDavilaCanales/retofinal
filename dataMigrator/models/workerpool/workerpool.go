package models

import "sync"

type WorkerPool[Input any, Output any] struct {
	taskExecutor TaskExecutor[Input, Output]
	taskQueueCh  chan Input
	resultCh     chan Result[Output]
	workerCount  int
	wg           *sync.WaitGroup
}

// Inicializa un WorkerPool
// Los parametros que necesita son:
//
//	taskExecutor: Es la funcion que cada worker ejecutara para realizar las tareas
//	taskQueueCh: Canal por el cual se enviaran las tareas a realizar
//	resultCh: Canal por el cual se enviaran los resultados de realizar las tareas
//	workerCount: Cantidad de workers que se crearan
// 	wg: WaitGroup para esperar a que todos los workers terminen
func NewWorkerPool[Input any, Output any](taskExecutor TaskExecutor[Input, Output], taskQueueCh chan Input, resultCh chan Result[Output], workerCount int, wg *sync.WaitGroup) *WorkerPool[Input, Output] {
	return &WorkerPool[Input, Output]{
		taskExecutor: taskExecutor,
		taskQueueCh:  taskQueueCh,
		resultCh:     resultCh,
		workerCount:  workerCount,
		wg:           wg,
	}
}

// Inicializa los workers
func (wp *WorkerPool[Input, Output]) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		worker := NewWorker(i+1, wp.taskExecutor, wp.taskQueueCh, wp.resultCh, wp.wg)
		worker.Start()
	}

	go func() {
		wp.wg.Wait()
		if wp.resultCh != nil {
			close(wp.resultCh)
		}
	}()
}
