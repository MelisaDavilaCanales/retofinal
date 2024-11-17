package models

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/enums"
	"gorm.io/gorm"
)

type Persons []Person

type Person struct {
	gorm.Model
	ID            uint                `json:"id" gorm:"primaryKey;autoIncrement"`
	Age           int                 `json:"age" gorm:"not null"`
	Workclass     enums.Workclass     `json:"workclass" gorm:"type:workclass"`
	Fnlwgt        int                 `json:"fnlwgt" gorm:"not null"`
	Education     string              `json:"education" gorm:"not null"`
	EducationNum  int                 `json:"education_num" gorm:"not null"`
	MaritalStatus enums.MaritalStatus `json:"marital_status" gorm:"type:marital_status;not null"`
	Occupation    enums.Occupation    `json:"occupation" gorm:"type:occupation"`
	Relationship  enums.Relationship  `json:"relationship" gorm:"type:relationship;not null"`
	Race          enums.Race          `json:"race" gorm:"type:race;not null"`
	Sex           enums.Sex           `json:"sex" gorm:"type:sex;not null"`
	CapitalGain   int                 `json:"capital_gain" gorm:"not null"`
	CapitalLoss   int                 `json:"capital_loss" gorm:"not null"`
	HoursPerWeek  int                 `json:"hours_per_week" gorm:"not null"`
	NativeCountry string              `json:"native_country"`
	Income        enums.Income        `json:"income" gorm:"type:income;not null"`
}
