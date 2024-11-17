package models

type Result[Output any] struct {
	Value Output
	Err   error
}
