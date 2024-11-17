package models

type TaskExecutor[Input any, Output any] func(workerIndex int, input Input) (result Output, err error)
