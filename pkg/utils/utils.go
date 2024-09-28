package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/models"
)

func GetDbConfigFromForm(c *gin.Context) models.DbConfig {
	return models.DbConfig{
		DbHost:     c.Request.FormValue("DbHost"),
		DbUser:     c.Request.FormValue("DbUser"),
		DbPort:     c.Request.FormValue("DbPort"),
		DbPassword: c.Request.FormValue("DbPassword"),
		DbName:     c.Request.FormValue("DbName"),
		DbSslMode:  c.Request.FormValue("DbSslMode"),
	}
}
