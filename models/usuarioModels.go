package models

type Usuario struct {
	Id_usuario       uint `gorm:"primaryKey" json:"id_usuario"`
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
