package models

type Usuario struct {
	Id_usuario       uint `gorm:"primary_Key"`
	Nombre           string
	Apellido         string
	Telefono         string
	Correo           string
	Fecha_nacimiento string
	Base
}

func (Usuario) TableName() string {
	return "usuario"
}
