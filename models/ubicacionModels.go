package models

import (
	"gorm.io/gorm"
)

type Ubicacion struct {
	gorm.Model
	Id_ubicacion uint
	Longitud     string
	Latitud      string
	Presicion    uint
}

func (Ubicacion) TableName() string {
	return "ubicacion"
}
