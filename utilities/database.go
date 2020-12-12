package utilities

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Database handle
var Database *sqlx.DB

// DB returns a handle to the database
func DB() *sqlx.DB {
	return Database
}

// ConfigureDB configure the database
func ConfigureDB() {

	// load environment variables from .env file
	loadEnvVariables()

	env := os.Getenv("ENVIRONMENT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	log.Infoln("Loaded environment variables for.....", env)

	//db, err := sqlx.Open("sqlite3", "./inventory.db")
	connectString := dbUser + ":" + dbPassword + "@(" + dbHost + ")/" +
		dbDatabase + "?charset=utf8&parseTime=True&loc=Local"
	log.Debugln("connectString ", connectString)

	db, err := sqlx.Connect("mysql", connectString)

	Database = db

	if err != nil {
		log.Panicln("Error establishing connection to database: ", err)
	}

	log.Infoln("Established connection to database ... ", dbHost, dbDatabase, db)

}

// ConfigureLogger configures app wide logger
func ConfigureLogger() {

	//example found here: https://stackoverflow.com/questions/48971780/change-format-of-log-output-logrus/48972299

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

	// print calling method in the log
	//log.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
	log.SetLevel(log.DebugLevel)

	log.Infoln("Logger configured for the app")

}

// load environment variables from file .env which must be in root directory
func loadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading environment variables...", err)
	}

}
