package db

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/parthvinchhi/db-backup/pkg/mock-db"
	"github.com/stretchr/testify/assert"
	// Adjust the path to your db package
)

func TestConnectPostgreSQL(t *testing.T) {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock PostgresDB
	mockDB := mockdb.NewMockPostgresDB(ctrl)

	// Set expectations
	mockDB.EXPECT().ConnectPostgreSQL().Return(nil).Times(1)

	// Call the method you are testing
	err := mockDB.ConnectPostgreSQL()

	// Validate the results
	assert.NoError(t, err)
}

func TestConnectPostgreSQL_Error(t *testing.T) {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock PostgresDB
	mockDB := mockdb.NewMockPostgresDB(ctrl)

	// Set expectations
	mockDB.EXPECT().ConnectPostgreSQL().Return(errors.New("failed to connect")).Times(1)

	// Call the method you are testing
	err := mockDB.ConnectPostgreSQL()

	// Validate the results
	assert.Error(t, err)
	assert.Equal(t, "failed to connect", err.Error())
}

func TestBackUpPostgreSQLData(t *testing.T) {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock PostgresDB
	mockDB := mockdb.NewMockPostgresDB(ctrl)

	// Expect the BackUpPostgreSQLData method to be called and return no error
	mockDB.EXPECT().BackUpPostgreSQLData().Return(nil).Times(1)

	// Call the method
	err := mockDB.BackUpPostgreSQLData()

	// Verify the result
	assert.NoError(t, err)
}

func TestBackUpPostgreSQLData_Error(t *testing.T) {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock PostgresDB
	mockDB := mockdb.NewMockPostgresDB(ctrl)

	// Expect the BackUpPostgreSQLData method to return an error
	mockDB.EXPECT().BackUpPostgreSQLData().Return(errors.New("backup failed")).Times(1)

	// Call the method
	err := mockDB.BackUpPostgreSQLData()

	// Verify the result
	assert.Error(t, err)
	assert.Equal(t, "backup failed", err.Error())
}

func TestRestorePostgreSQLData(t *testing.T) {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock PostgresDB
	mockDB := mockdb.NewMockPostgresDB(ctrl)

	// Create a temporary file for simulating the SQL backup file
	tmpfile, err := os.CreateTemp("", "backup.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Expect the RestorePostgreSQLData method to be called and return no error
	mockDB.EXPECT().RestorePostgreSQLData(tmpfile.Name()).Return(nil).Times(1)

	// Call the method
	err = mockDB.RestorePostgreSQLData(tmpfile.Name())

	// Verify the result
	assert.NoError(t, err)
}

func TestRestorePostgreSQLData_Error(t *testing.T) {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock PostgresDB
	mockDB := mockdb.NewMockPostgresDB(ctrl)

	// Expect the RestorePostgreSQLData method to return an error
	mockDB.EXPECT().RestorePostgreSQLData("invalid_file_path").Return(errors.New("file not found")).Times(1)

	// Call the method
	err := mockDB.RestorePostgreSQLData("invalid_file_path")

	// Verify the result
	assert.Error(t, err)
	assert.Equal(t, "file not found", err.Error())
}
