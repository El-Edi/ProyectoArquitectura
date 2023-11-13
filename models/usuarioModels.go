package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Id_usuario       uint
	Nombre           string
	Apellido         string
	Telefono         string
	Correo           string
	Fecha_nacimiento string
}

func (Usuario) TableName() string {
	return "usuario"
}
