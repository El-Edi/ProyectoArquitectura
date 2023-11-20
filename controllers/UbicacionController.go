package controllers

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"

	"github.com/gin-gonic/gin"
)

func UbicacionGet(contextServe *gin.Context) {
	var ubicacion []models.Ubicacion
	configs.DB.Limit(100).Find(&ubicacion)
	contextServe.JSON(200, &ubicacion)
	return
}

func UbicacionGetById(contextServe *gin.Context) {
	id_usuario := contextServe.Param("id")
	var ubicacion models.Ubicacion
	configs.DB.Where("id_usuario = ?", id_usuario).Find(&ubicacion)
	contextServe.JSON(200, &ubicacion)
	return
}
