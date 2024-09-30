package db

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/parthvinchhi/db-backup/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Helper function to mock exec.Command
func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

// TestHelperProcess is a mock exec command handler
func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Println("mocked command")
	os.Exit(0)
}

func TestConnectPostgreSQL(t *testing.T) {
	// Mock the database connection and ping
	datab, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer datab.Close()

	// Mock GORM DB connection with sqlmock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: datab,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Create a mock instance of Postgres
	mockPostgres := &Postgres{
		Config: models.DbConfig{
			DbHost:     "localhost",
			DbPort:     "5432",
			DbUser:     "test_user",
			DbPassword: "test_pass",
			DbName:     "test_db",
			DbSslMode:  "disable",
		},
		db: gormDB,
	}

	// Expect ping
	mock.ExpectPing()

	// Test the connection
	err = mockPostgres.ConnectPostgreSQL()
	assert.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestPostgres_BackUpPostgreSQLData(t *testing.T) {
	// Mock the exec.Command
	execCommand := exec.Command
	defer func() { exec.Command = execCommand }()

	exec.Command = func(name string, args ...string) *exec.Cmd {
		// Simulate successful command execution
		return &exec.Cmd{}
	}

	// Mock Postgres instance with DbConfig and Helper
	mockPostgres := &Postgres{
		Config: models.DbConfig{
			DbHost:     "localhost",
			DbPort:     "5432",
			DbUser:     "test_user",
			DbPassword: "test_pass",
			DbName:     "test_db",
			DbSslMode:  "disable",
		},
		Helper: models.Helper{},
	}

	// Run the backup method
	err := mockPostgres.BackUpPostgreSQLData()
	assert.NoError(t, err)
	assert.Contains(t, mockPostgres.Helper.BackupFile, "test_db_backup_")
}

func TestPostgres_RestorePostgreSQLData(t *testing.T) {
	// Create sqlmock
	datab, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer datab.Close()

	// Mock GORM DB connection with sqlmock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: datab,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Mock Postgres instance with DbConfig
	mockPostgres := &Postgres{
		Config: models.DbConfig{
			DbHost:     "localhost",
			DbPort:     "5432",
			DbUser:     "test_user",
			DbPassword: "test_pass",
			DbName:     "test_db",
			DbSslMode:  "disable",
		},
		db: gormDB,
	}

	// Mock file opening
	fileContent := "some SQL content"
	file, err := os.CreateTemp("", "restore_test.sql")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.WriteString(fileContent)
	assert.NoError(t, err)

	// Test the restore method
	err = mockPostgres.RestorePostgreSQLData(file.Name())
	assert.NoError(t, err)

	// Expect Exec to be called with the file content
	mock.ExpectExec("some SQL content").WillReturnResult(sqlmock.NewResult(1, 1))

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
