package controllers

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"

	"github.com/gin-gonic/gin"
)

type EventoRequestBody struct {
	Id_evento    uint   `json:"id_evento"`
	Tipo         string `json:"tipo"`
	Descripcion  string `json:"descripcion"`
	Id_ubicacion uint   `json:"id_ubicacion"`
}

func EventoCreate(contextServe *gin.Context) {

	body := EventoRequestBody{}

	contextServe.BindJSON(&body)

	evento := &models.Evento{Id_evento: body.Id_evento, Tipo: body.Tipo, Descripcion: body.Descripcion, Id_ubicacion: body.Id_ubicacion}

	result := configs.DB.Create(&evento)

	if result.Error != nil {
		contextServe.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	contextServe.JSON(200, &evento)
}

func EventoGet(contextServe *gin.Context) {
	var evento []models.Evento
	configs.DB.Limit(100).Find(&evento)
	contextServe.JSON(200, &evento)
	return
}
