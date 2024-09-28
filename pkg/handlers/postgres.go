package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/db"
	"github.com/parthvinchhi/db-backup/pkg/utils"
)

func BackupPostgreSQLHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)

	// Validate if all required fields are provided
	if dbConfig.DbHost == "" || dbConfig.DbUser == "" || dbConfig.DbPort == "" || dbConfig.DbPassword == "" || dbConfig.DbName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required database details"})
		return
	}

	// Call the function to perform backup or any other operation
	postgres := db.Postgres{
		Config: dbConfig,
	}

	if err := postgres.ConnectPostgreSQL(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := postgres.BackUpPostgreSQLData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "Backup successful"})
}

func RestorePostgreSQLHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)
	backupFile := c.Request.FormValue("backup_file") // Fetching backup file path

	// Validate if all required fields are provided
	if dbConfig.DbHost == "" || dbConfig.DbUser == "" || dbConfig.DbPort == "" || dbConfig.DbPassword == "" || dbConfig.DbName == "" || backupFile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required details"})
		return
	}

	// Call the function to perform restore
	postgres := db.Postgres{
		Config: dbConfig,
	}

	if err := postgres.ConnectPostgreSQL(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := postgres.RestorePostgreSQLData(backupFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "Restore successful"})
}
