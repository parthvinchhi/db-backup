package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/handlers"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.POST("/postgres/backup", handlers.BackupPostgreSQLHandler)
	r.POST("/postgres/restore", handlers.RestorePostgreSQLHandler)

	return r
}
