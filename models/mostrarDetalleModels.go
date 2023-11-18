package models

import (
	"gorm.io/gorm"
)

type MostrarDetalle struct {
	gorm.Model
	Id_Detalle uint
	Id_evento string
	descripcion string
}

func (MostrarDetalle) TableName() string {
	return "MostrarDetalle"
}
