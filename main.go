package main

import (

	/*
		The packages below are my packages.  I need to give a full path under the src directory.  So config and models are under github.com/sodabitsdev/inventory-app/
	*/

	/*
		database/sql

		https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html

		SQL tutorial in GO
		https://www.alexedwards.net/blog/practical-persistence-sql
	*/

	// "github.com/sodabitsdev/inventory-app/config"

	//"time"

	//"github.com/joho/godotenv"
	"github.com/sodabitsdev/inventory-app/routes"
	"github.com/sodabitsdev/inventory-app/utilities"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

// main function to start the server
func main() {

	// configure logger
	utilities.ConfigureLogger()

	// configure database
	utilities.ConfigureDB()

	// env := os.Getenv("ENVIRONMENT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbDatabase := os.Getenv("DB_DATABASE")
	// log.Infoln("Loaded environment variables for.....", env)

	// //db, err := sqlx.Open("sqlite3", "./inventory.db")
	// connectString := dbUser + ":" + dbPassword + "@(" + dbHost + ")/" + dbDatabase
	// log.Debugln("connectString ", connectString)

	// db, err := sqlx.Connect("mysql", connectString)

	// if err != nil {
	// 	log.Panicln("Error establishing connection to database: ", err)
	// }

	// defer DB.Close()

	// log.Infoln("Established connection to database ... ", dbHost, dbDatabase, db)

	// configure HTTP
	router := routes.SetupRouter()

	// run HTTP
	router.Run()

}

// // configureLogger configures app wide logger
// func configureLogger() {

// 	//example found here: https://stackoverflow.com/questions/48971780/change-format-of-log-output-logrus/48972299

// 	log.SetFormatter(&log.TextFormatter{
// 		DisableColors:   false,
// 		FullTimestamp:   true,
// 		TimestampFormat: "2006-01-02 15:04:05",
// 		ForceColors:     true,
// 	})

// 	// print calling method in the log
// 	//log.SetReportCaller(true)

// 	// Output to stdout instead of the default stderr
// 	// Can be any io.Writer, see below for File example
// 	log.SetOutput(os.Stdout)

// 	// Only log the warning severity or above.
// 	//log.SetLevel(log.WarnLevel)
// 	log.SetLevel(log.DebugLevel)

// }
