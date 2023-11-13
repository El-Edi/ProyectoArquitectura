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

func UsuarioCreate(c *gin.Context) {

	body := UsuarioRequestBody{}

	c.BindJSON(&body)

	usuario := &models.Usuario{Id_usuario: body.Id_usuario, Nombre: body.Nombre, Apellido: body.Apellido, Telefono: body.Telefono, Correo: body.Correo, Fecha_nacimiento: body.Fecha_nacimiento}

	result := configs.DB.Create(&usuario)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &usuario)
}

func UsuarioGet(c *gin.Context) {
	var usuario []models.Usuario
	configs.DB.Limit(100).Find(&usuario)
	c.JSON(200, &usuario)
	return
}
