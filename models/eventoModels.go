package models

import (
	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
	Id_evento    string 
	Tipo         string 
	Descripcion  string 
	Id_ubicacion uint   
}

func (Evento) TableName() string {
	return "evento"
}