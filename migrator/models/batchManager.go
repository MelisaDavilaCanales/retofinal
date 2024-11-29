package models

import "github.com/MelisaDavilaCanales/trucode3-challenge-final.git/shared"

type BatchManager struct {
	BatchCount int
	Batches    []PersonBatch
	BatchSize  int
}

func NewBatchManager(batchCount int) *BatchManager {
	return &BatchManager{
		Batches:    make([]PersonBatch, batchCount),
		BatchSize:  shared.BatchSize,
		BatchCount: batchCount,
	}
}

func (bm *BatchManager) InitBatches() {
	for i := 0; i < bm.BatchCount; i++ {
		bm.Batches[i] = PersonBatch{
			IdBatch:   i + 1,
			Persons:   make([]Person, 0),
			BatchSize: bm.BatchSize,
		}
	}
}

func (bm *BatchManager) GetBatch(id int) (*PersonBatch, bool) {
	for i := range bm.Batches {
		if bm.Batches[i].IdBatch == id {
			return &bm.Batches[i], true
		}
	}
	return nil, false
}
