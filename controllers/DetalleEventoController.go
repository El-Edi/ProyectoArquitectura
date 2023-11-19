package controllers

import (
	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func DetalleEventoGetById(c *gin.Context) {
	fmt.Print("Entro")
	Id_evento := c.Param("id")
	var detalle models.DetalleEvento
	configs.DB.Where("Id_evento = ?", Id_evento).Find(&detalle)
	c.JSON(200, &detalle)
	return
}
