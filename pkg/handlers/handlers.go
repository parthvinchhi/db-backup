package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/utils"
)

func BackupDataHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)

	switch dbConfig.DbType {
	case "PostgreSQL":
		if !validateDbConfig(&dbConfig) {
			handleErrors(c, http.StatusBadRequest, "Missing required database details")
			return
		}

		postgres, err := getDbConnection(dbConfig, c)
		if err != nil {
			handleErrors(c, http.StatusInternalServerError, "Failed to connect to database")
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Minute)
		defer cancel()

		if err := postgres.BackUpPostgreSQLData(ctx); err != nil {
			handleErrors(c, http.StatusInternalServerError, "Failed to backup database")
			return
		}

		// Send a success response
		c.JSON(http.StatusOK, gin.H{"message": "Backup successful"})

	case "MySQL":
	case "MongoDB":
	default:
	}
}

func RestoreDataHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)
	backupFile := c.Request.FormValue("backup_file") // Fetching backup file path

	switch dbConfig.DbType {
	case "PostgreSQL":
	case "MySQL":
	case "MongoDB":
	default:
	}
}
