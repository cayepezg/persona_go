package utils

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type conf struct {
	ApiServerPort  string
	DBHost         string
	DBPort         int
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	DBMaxIdleConns int
	DBMaxOpenConns int
}

var (
	config conf
	oc     sync.Once
)

// InitConf inicializa la configuración de la aplicación.
func InitConf() {

	oc.Do(func() {
		// Puerto que escuchará el API
		if os.Getenv("API_SERVER_PORT") != "" {
			config.ApiServerPort = os.Getenv("API_SERVER_PORT")
		} else {
			config.ApiServerPort = "80"
		}

		if os.Getenv("DB_HOST") != "" {
			config.DBHost = os.Getenv("DB_HOST")
		} else {
			config.DBHost = "localhost"
		}

		if os.Getenv("DB_PORT") != "" {
			dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
			if err != nil {
				log.Println("No se suministró un puerto de base de datos válido. Se usará el valor por defecto (5432)")
				config.DBPort = 5432
			} else {
				config.DBPort = dbPort
			}
		}

		if os.Getenv("DB_USER") != "" {
			config.DBUser = os.Getenv("DB_USER")
		} else {
			config.DBUser = "user"
		}

		if os.Getenv("DB_PASSWORD") != "" {
			config.DBPassword = os.Getenv("DB_PASSWORD")
		} else {
			config.DBPassword = "password"
		}

		if os.Getenv("DB_NAME") != "" {
			config.DBName = os.Getenv("DB_NAME")
		} else {
			config.DBName = "knoios"
		}

		if os.Getenv("DB_SSL_MODE") != "" {
			config.DBSSLMode = os.Getenv("DB_SSL_MODE")
		} else {
			config.DBSSLMode = "disable"
		}

		if os.Getenv("DB_MAXIDLECONNS") != "" {
			db_maxidleconns, err := strconv.Atoi(os.Getenv("DB_MAXIDLECONNS"))
			if err == nil {
				config.DBMaxIdleConns = db_maxidleconns
			} else {
				log.Printf("Error parseando DB_MAXIDLECONNS = %s . %s .Se usará el valor 1.\n",
					os.Getenv("DB_MAXIDLECONNS"), err.Error())
				config.DBMaxIdleConns = 1
			}
		} else {
			config.DBMaxIdleConns = 1
		}

		if os.Getenv("DB_MAXOPENCONNS") != "" {
			db_maxopenconns, err := strconv.Atoi(os.Getenv("DB_MAXOPENCONNS"))
			if err == nil {
				config.DBMaxOpenConns = db_maxopenconns
			} else {
				log.Printf("Error parseando DB_MAXOPENCONNS = %s . %s .Se usará el valor 1.\n",
					os.Getenv("DB_MAXOPENCONNS"), err.Error())
				config.DBMaxOpenConns = 1
			}
		} else {
			config.DBMaxOpenConns = 1
		}

	})
}

func GetConfig() conf {
	return config
}
