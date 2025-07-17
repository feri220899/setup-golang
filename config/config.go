package Config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	mysql "gorm.io/driver/mysql"
	postgres "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // <-- tambahkan ini
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}

func ConnectDB() *gorm.DB {
	InitConfig()

	driver := viper.GetString("DB_DRIVER")
	var db *gorm.DB
	var err error

	switch driver {
	case "mysql":
		user := viper.GetString("DB_USER")
		pass := viper.GetString("DB_PASS")
		host := viper.GetString("DB_HOST")
		port := viper.GetString("DB_PORT")
		name := viper.GetString("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, name)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "postgres":
		user := viper.GetString("DB_USER")
		pass := viper.GetString("DB_PASS")
		host := viper.GetString("DB_HOST")
		port := viper.GetString("DB_PORT")
		name := viper.GetString("DB_NAME")

		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			host, port, user, name, pass)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	case "sqlite":
		path := viper.GetString("DB_SQLITE_PATH")
		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	default:
		log.Fatalf("Database driver %s tidak didukung", driver)
	}

	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	return db
}
