package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/db-backup/pkg/db"
	"github.com/parthvinchhi/db-backup/pkg/models"
)

func BackupPostgreSQLHandler(c *gin.Context) {
	var request models.DbConfig

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pg := &db.Postgres{
		Config: request,
		Helper: models.Helper{},
	}

	if err := pg.ConnectPostgreSQL(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := pg.BackUpPostgreSQLData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Backup completed successfully"})
}

func RestorePostgreSQLHandler(c *gin.Context) {
	var request struct {
		models.DbConfig
		FilePath string `json:"file_path" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initialize Postgres instance
	pg := &db.Postgres{
		Config: request.DbConfig,
		Helper: models.Helper{}, // Add necessary initialization
	}

	// Connect to the PostgreSQL database
	if err := pg.ConnectPostgreSQL(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Perform restore
	if err := pg.RestorePostgreSQLData(request.FilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restore completed successfully"})
}
