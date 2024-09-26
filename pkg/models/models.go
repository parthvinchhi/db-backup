package models

type DbConfig struct {
	DbHost     string `json:"db_host"`
	DbUser     string `json:"db_user"`
	DbPort     string `json:"db_port"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
	DbSslMode  string `json:"db_sslmode"`
}

type Helper struct {
	TimeStamp   string `json:"time_stamp"`
	BackupFile  string `json:"backup_file"`
	RestoreFile string `json:"restore_file"`
}
