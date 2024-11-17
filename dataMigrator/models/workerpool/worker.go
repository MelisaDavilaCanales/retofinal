package models

import "sync"

type Worker[Input any, Output any] struct {
	Id           int
	taskExecutor TaskExecutor[Input, Output]
	taskQueueCh  chan Input
	resultCh     chan Result[Output]
	wg           *sync.WaitGroup
}

// Inicializa un Worker
// Los parametros que necesita son:
//
// 	id: identificador del worker
//	taskExecutor: funcion que ejecutara la tareas
//	taskQueueCh: canal del que estara escuchando y tomara las tareas a realizar
//	resultCh: canal por el cual escribira el resultado retornado por taskExecutor
//	wg: WaitGroup para decrementar el contador cunado el worker termine
func NewWorker[Input any, Output any](id int, taskExecutor TaskExecutor[Input, Output], tasksCh chan Input, resultCh chan Result[Output], wg *sync.WaitGroup) *Worker[Input, Output] {
	return &Worker[Input, Output]{
		Id:           id,
		taskExecutor: taskExecutor,
		taskQueueCh:  tasksCh,
		resultCh:     resultCh,
		wg:           wg,
	}
}

// El wprker comienza a escuchar el canal de tareas y ejecutarlas
func (w *Worker[Input, Output]) Start() {
	go func() {
		defer w.wg.Done()
		for task := range w.taskQueueCh {
			result, err := w.taskExecutor(w.Id, task)
			if w.resultCh != nil && !isNil(result) {
				w.resultCh <- Result[Output]{Value: result, Err: err}
			}
		}
	}()
}
