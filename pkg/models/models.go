package models

type DbConfig struct {
	DbType     string
	DbHost     string
	DbUser     string
	DbPort     string
	DbPassword string
	DbName     string
	DbSslMode  string
}

type Helper struct {
	TimeStamp   string
	BackupFile  string
	RestoreFile string
}
