package controllers

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"

	"github.com/gin-gonic/gin"
)

type UsuarioRequestBody struct {
	Id_usuario       uint   `json:"id_usuario"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Telefono         string `json:"telefono"`
	Correo           string `json:"correo"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}

func UsuarioCreate(contextServe *gin.Context) {

	body := UsuarioRequestBody{}

	contextServe.BindJSON(&body)

	usuario := &models.Usuario{Id_usuario: body.Id_usuario, Nombre: body.Nombre, Apellido: body.Apellido, Telefono: body.Telefono, Correo: body.Correo, Fecha_nacimiento: body.Fecha_nacimiento}

	result := configs.DB.Create(&usuario)

	if result.Error != nil {
		contextServe.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	contextServe.JSON(200, &usuario)
}

func UsuarioGet(contextServe *gin.Context) {
	var usuario []models.Usuario
	configs.DB.Limit(100).Find(&usuario)
	contextServe.JSON(200, &usuario)
	return
}
