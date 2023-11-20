package models

type DetalleEvento struct {
	Id_Detalle  uint `gorm:"primary_Key"`
	Descripcion string
	Id_evento   uint
	// Id_evento   uint   `gorm:"column:id_evento;index" json:"custom_DetalleEvento_id"`
	// Evento      Evento `gorm:"foreignKey:Id_evento" json:"author"`
	Base
}

func (DetalleEvento) TableName() string {
	return "DetalleEvento"
}
