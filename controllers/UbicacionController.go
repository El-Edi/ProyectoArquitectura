package controllers

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"

	"github.com/gin-gonic/gin"
)

func UbicacionGet(c *gin.Context) {
	var ubicacion []models.Ubicacion
	configs.DB.Limit(100).Find(&ubicacion)
	c.JSON(200, &ubicacion)
	return
}

func UbicacionGetById(c *gin.Context) {
	id_usuario := c.Param("id")
	var ubicacion models.Ubicacion
	configs.DB.Where("id_usuario = ?", id_usuario).Find(&ubicacion)
	c.JSON(200, &ubicacion)
	return
}
