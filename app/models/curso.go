package models

import (
	"fmt"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Curso struct {
	Id         string `gorm:"primaryKey;"`
	Nombre    string
	Area     string
	Cargas []Carga_acad
}

func (tab Curso) ToString() string {
	return tab.Nombre
}

func (tab *Curso) BeforeCreate(*gorm.DB) error {
	tab.Id = uuid.NewV4().String()
	return nil
}

func (curso Curso) FindAll(conn *gorm.DB) ([]Curso, error) {
	var cursos []Curso
	if err := conn.Preload("Cargas").Find(&cursos).Error; err != nil {
		return nil, err
	}
	return cursos, nil
}

func (curso Curso) GetAll(conn *gorm.DB) ([]Curso, error) {
	var cursos []Curso
	if err := conn.Find(&cursos).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//fmt.Printf("Error: %v", err)
		//return fmt.Errorf("Error: %v", err)
		//continue
		return nil, fmt.Errorf("Error: %v", err)
	}
	return cursos, nil
}