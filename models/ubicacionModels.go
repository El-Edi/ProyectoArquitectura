package models

type Ubicacion struct {
	Id_ubicacion uint
	Longitud     string
	Latitud      string
	Precision    uint
	Id_usuario   uint    `gorm:"column:id_usuario;index" json:"custom_usuario_id"`
	Usuario      Usuario `gorm:"foreignKey:Id_usuario" json:"author"`
	Base
}

func (Ubicacion) TableName() string {
	return "ubicacion"
}
