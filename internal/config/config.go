package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/sailsforce/gomicro-kit/utils"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Psql *gorm.DB
var Logger *logrus.Logger

// NewServiceConfig : function that inits the database, logger, and other parts of the config.
// env variables needed
// LOG_LEVEL, DATABASE_URL
func NewServiceConfig(withDb bool) error {
	log.Println("Creating Togo Read Micro-service config...")
	log.Println("init logger...")
	// create logger using sirupsen/Logrus
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	Logger.SetOutput(os.Stdout)
	logLvl, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		Logger.Info("using default log level")
		logLvl = 4
	}
	Logger.SetLevel(logLvl)
	log.Println("logger done.")
	// logger setup done.

	// GORM setup
	log.Println("init Gorm DB...")
	if withDb {
		dbUrl := os.Getenv("DATABASE_URL")
		Logger.Info("connecting to db: ", utils.GetDSN(dbUrl))
		db, err := sql.Open("postgres", utils.GetDSN(dbUrl))
		if err != nil {
			Logger.Error("sql connection error")
			return err
		}
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(logLvl)),
		})
		if err != nil {
			Logger.Error("gorm connection error")
			return err
		}
		Psql = gormDB
		// GORM setup done.
	}
	log.Println("Gorm DB done.")
	Logger.Info("config created.")

	return nil
}
