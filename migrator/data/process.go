package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/enums"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/shared"
)

var BatchManager *models.BatchManager

func init() {
	BatchManager = models.NewBatchManager(shared.ProcessingWorkersCount)
	BatchManager.InitBatches()
}

func ProcessDataPersons(idWorker int, data string) (models.Persons, error) {
	person := parsePerson(data)
	batch, isExist := BatchManager.GetBatch(idWorker)
	if !isExist {
		return nil, fmt.Errorf("no existe el batch con id %d", idWorker)
	}

	batchPerson, isBatchFull := BuildBatchOfPersons(person, batch)
	if isBatchFull {
		tx := database.DBConn.CreateInBatches(batchPerson.Persons, shared.BatchSize)
		if tx.Error != nil {
			return nil, fmt.Errorf("error al insertar en la base de datos: %w", tx.Error)
		}
		return batchPerson.Persons, nil
	}
	return nil, nil
}

func BuildBatchOfPersons(person models.Person, batch *models.PersonBatch) (models.PersonBatch, bool) {
	batch.Persons = append(batch.Persons, person)
	if len(batch.Persons) == batch.BatchSize {
		batchTemp := *batch
		batch.ResetBatch()
		return batchTemp, true
	}
	return models.PersonBatch{}, false
}

func parsePerson(data string) models.Person {
	fields := strings.Split(data, ", ")
	if len(fields) < 15 {
		fmt.Println("Incomplete Data")
		return models.Person{}
	}
	age := toInt(fields[0])
	workclass := enums.Workclass(ifIsNull(fields[1]))
	fnlwgt := toInt(fields[2])
	education := fields[3]
	educationNum := toInt(fields[4])
	maritalStatus := enums.MaritalStatus(fields[5])
	occupation := enums.Occupation(ifIsNull(fields[6]))
	relationship := enums.Relationship(fields[7])
	race := enums.Race(fields[8])
	sex := enums.Sex(fields[9])
	capitalGain := toInt(fields[10])
	capitalLoss := toInt(fields[11])
	hoursPerWeek := toInt(fields[12])
	nativeCountry := ifIsNull(fields[13])
	income := enums.Income(fields[14])
	person := models.Person{
		Age:           age,
		Workclass:     workclass,
		Fnlwgt:        fnlwgt,
		Education:     education,
		EducationNum:  educationNum,
		MaritalStatus: maritalStatus,
		Occupation:    occupation,
		Relationship:  relationship,
		Race:          race,
		Sex:           sex,
		CapitalGain:   capitalGain,
		CapitalLoss:   capitalLoss,
		HoursPerWeek:  hoursPerWeek,
		NativeCountry: nativeCountry,
		Income:        income,
	}
	return person
}

func toInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

func ifIsNull(value string) string {
	if value == "?" {
		return ""
	}
	return value
}
