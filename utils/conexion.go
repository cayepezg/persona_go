package utils

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
	err  error
)

// GetConnection Obtiene una conexión a base de datos
func GetConnection() *gorm.DB {

	once.Do(func() {
		db, err = getConnection()
		if err != nil || db == nil {
			log.Fatal("No puedo conectarme a la base de datos")
		}
	})
	return db
}

func getConnection() (*gorm.DB, error) {

	conf := GetConfig()

	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBSSLMode)
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	log.Println(dbURL)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Error grave durante el intento de conexión a base de datos. Por favor, verifique la configuración que está utilizando")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(conf.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.DBMaxOpenConns)

	err = sqlDB.Ping()

	if err != nil {
		log.Println("Error de conexión a base de datos. ", err.Error())
	} else {
		log.Println("Conexión a Base de datos realizada exitosamente")
	}

	return db, err
}

// TestConnection Prueba conexión a base de datos.
func TestConnection() {
	GetConnection()
}
