package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Carga_acad struct {
	Id string `gorm:"primary_key;"`
	//Fecha  string
	Semestre string
	CursoId string `gorm:"size:191"`
	Curso   Curso //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"embedded"` crea el compo nombres y codigo de alumnos
}

func (tab Carga_acad) ToString() string {
	return fmt.Sprintf("id: %d\nSemestre: %s", tab.Id, tab.Semestre)
}

func (tab *Carga_acad) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

// Curso   Curso //para crear el FK `gorm:"foreignkey:CursoId"`
