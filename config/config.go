package config

import "fmt"

const (
	DBHost     = "localhost"
	DBPort     = "3306"
	DBUser     = "root"
	DBPAssword = "abc"
	DBName     = "listatareas"
)

func GetDBConnection() string {
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPAssword,
		DBHost,
		DBPort,
		DBName)
	return connection
}
