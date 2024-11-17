package data

import (
	"fmt"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	modelsWP "github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models/workerpool"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/shared"
)

func SendBathToBD(workerIndex int, batch modelsWP.Result[models.Persons]) (bool, error) {

	if batch.Err != nil {
		return false, batch.Err
	}
	for _, persons := range batch.Value {
		tx := database.DBConn.CreateInBatches(persons, shared.BatchSize)
		fmt.Println(" SENT TO DB ")
		if tx.Error != nil {
			return false, fmt.Errorf("error al insertar en la base de datos: %w", tx.Error)
		}
	}
	return true, nil
}
