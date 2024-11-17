package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/data"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	models_wp "github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models/workerpool"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/shared"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/utils"
)

func main() {

	timeInit := time.Now()
	utils.GetEnvVars()

	database.CreateDbConnection()
	err := database.DBConn.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatal("Error al migrar los modelos de la base de datos")
	}

	personStringCh := make(chan string, shared.BufferCapacity)
	personGroupCh := make(chan models_wp.Result[models.Persons], shared.BufferCapacity)
	var wgReadFile, wgProcessData, wgSendData sync.WaitGroup

	wgReadFile.Add(1)
	go data.ReadFile(personStringCh, &wgReadFile)

	DataProcessPool := models_wp.NewWorkerPool(data.ProcessDataPersons, personStringCh, personGroupCh, shared.ProcessingWorkersCount, &wgProcessData)
	DataProcessPool.Start()

	DataSendPool := models_wp.NewWorkerPool(data.SendBathToBD, personGroupCh, nil, shared.SendingWorkersCount, &wgSendData)
	DataSendPool.Start()

	wgSendData.Wait()
	timeSince := time.Since(timeInit)
	fmt.Println("Tiempo de ejecución: ", timeSince)

}
