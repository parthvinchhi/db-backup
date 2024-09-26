package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/db"
	"github.com/parthvinchhi/db-backup/pkg/models"
)

func BackupPostgreSQLHandler(c *gin.Context) {
	var request models.DbConfig
	request.DbHost = c.Request.FormValue("db_host")
	request.DbUser = c.Request.FormValue("db_user")
	request.DbPort = c.Request.FormValue("db_port")
	request.DbPassword = c.Request.FormValue("db_password")
	request.DbName = c.Request.FormValue("db_name")
	request.DbSslMode = c.Request.FormValue("db_ssl_mode")

	// Validate if all required fields are provided
	if request.DbHost == "" || request.DbUser == "" || request.DbPort == "" || request.DbPassword == "" || request.DbName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required database details"})
		return
	}

	// Call the function to perform backup or any other operation
	postgres := db.Postgres{
		Config: request,
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
	var request models.DbConfig
	request.DbHost = c.Request.FormValue("db_host")
	request.DbUser = c.Request.FormValue("db_user")
	request.DbPort = c.Request.FormValue("db_port")
	request.DbPassword = c.Request.FormValue("db_password")
	request.DbName = c.Request.FormValue("db_name")
	request.DbSslMode = c.Request.FormValue("db_ssl_mode")
	backupFile := c.Request.FormValue("backup_file") // Fetching backup file path

	// Validate if all required fields are provided
	if request.DbHost == "" || request.DbUser == "" || request.DbPort == "" || request.DbPassword == "" || request.DbName == "" || backupFile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required details"})
		return
	}

	// Call the function to perform restore
	postgres := db.Postgres{
		Config: request,
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
