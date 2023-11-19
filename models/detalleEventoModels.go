package models

import (
	"gorm.io/gorm"
)

type DetalleEvento struct {
	gorm.Model
	Id_Detalle uint
	Id_evento string
	descripcion string
}

func (DetalleEvento) TableName() string {
	return "DetalleEvento"
}
