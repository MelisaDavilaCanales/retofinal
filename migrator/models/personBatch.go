package models

type PersonBatch struct {
	IdBatch   int
	Persons   Persons
	BatchSize int
}

func (bm *PersonBatch) ResetBatch() {
	bm.Persons = make([]Person, 0)
}
