package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/db"
	"github.com/parthvinchhi/db-backup/pkg/utils"
)

func BackupMongoDBHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)

	if dbConfig.DbHost == "" || dbConfig.DbUser == "" || dbConfig.DbPort == "" || dbConfig.DbPassword == "" || dbConfig.DbName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required database details"})
		return
	}

	// Create MongoDB instance and perform backup
	mongo := db.MongoDB{
		Config: dbConfig,
	}

	if err := mongo.ConnectMongoDb(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mongo.BackUpMongoDBData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Backup successful"})
}

func RestoreMongoDBHandler(c *gin.Context) {
	dbConfig := utils.GetDbConfigFromForm(c)
	backupDir := c.Request.FormValue("backup_dir")

	if dbConfig.DbHost == "" || dbConfig.DbUser == "" || dbConfig.DbPort == "" || dbConfig.DbPassword == "" || dbConfig.DbName == "" || backupDir == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required details"})
		return
	}

	// Create MongoDB instance and restore data
	mongo := db.MongoDB{
		Config: dbConfig,
	}

	if err := mongo.ConnectMongoDb(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mongo.RestoreMongoDBData(backupDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restore successful"})
}
