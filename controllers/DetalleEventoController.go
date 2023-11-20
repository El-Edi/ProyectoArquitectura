package controllers

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"

	"github.com/gin-gonic/gin"
)

func DetalleEventoGetById(contextServe *gin.Context) {
	Id_evento := contextServe.Param("id")
	var detalle models.DetalleEvento
	configs.DB.Where("Id_evento = ?", Id_evento).Find(&detalle)
	contextServe.JSON(200, &detalle)
	return
}
