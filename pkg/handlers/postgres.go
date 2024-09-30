package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/db"
	"github.com/parthvinchhi/db-backup/pkg/models"
	"github.com/parthvinchhi/db-backup/pkg/utils"
)

func validateDbConfig(dbConfig *models.DbConfig) bool {
	return dbConfig.DbHost != "" || dbConfig.DbUser != "" || dbConfig.DbPort != "" ||
		dbConfig.DbPassword != "" || dbConfig.DbName != "" || dbConfig.DbSslMode != ""
}

func handleErrors(c *gin.Context, status int, errMsg string) {
	c.JSON(status, gin.H{"error": errMsg})
}

func getDbConnection(dbConfig models.DbConfig, c *gin.Context) (*db.Postgres, error) {
	postgres := &db.Postgres{
		Config: dbConfig,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := postgres.ConnectPostgreSQL(ctx); err != nil {
		return nil, err
	}

	return postgres, nil
}

func BackupPostgreSQLHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)

	// Validate if all required fields are provided
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
