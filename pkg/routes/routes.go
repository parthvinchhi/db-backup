package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/handlers"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.POST("/postgres/backup", handlers.BackupPostgreSQLHandler)
	r.POST("/postgres/restore", handlers.RestorePostgreSQLHandler)
	r.POST("/mysql/backup", handlers.BackupMySQLHandler)
	r.POST("/mysql/restore", handlers.RestoreMySQLHandler)
	r.POST("/mongodb/backup", handlers.BackupMongoDBHandler)
	r.POST("/mongodb/restore", handlers.RestoreMongoDBHandler)

	return r
}
