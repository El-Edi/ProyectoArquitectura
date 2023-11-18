package main

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Usuario{}, &models.Ubicacion{},&models.MostrarDetalle{})
}
