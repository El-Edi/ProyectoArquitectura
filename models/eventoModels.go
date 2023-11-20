package models

type Evento struct {
	Id_evento    uint `gorm:"primary_Key"`
	Tipo         string
	Descripcion  string
	Id_ubicacion uint
	// Id_ubicacion uint      `gorm:"column:id_ubicacion;index" json:"custom_ubicacion_id"`
	// Ubicacion    Ubicacion `gorm:"foreignKey:Id_ubicacion" json:"author"`
	Base
}

func (Evento) TableName() string {
	return "evento"
}
