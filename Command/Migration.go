package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm-training/Database/Migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type DatabaseConnection struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBSchema   string
}

func main() {
	// Load .env to command
	errGoDotEnv := godotenv.Load()
	if errGoDotEnv != nil {
		log.Panic("Error load .env file, please refers this stack trace for debugging: " + errGoDotEnv.Error())
	}

	// Get variable in .env
	appEnvironment := os.Getenv("APP_ENVIRONMENT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_SCHEMA")

	// Insert database variable to DBConnection Struct
	databaseConnection := &DatabaseConnection{
		DBHost:     dbHost,
		DBUsername: dbUsername,
		DBPassword: dbPassword,
		DBSchema:   dbSchema,
	}

	// Validate app environment
	checkAppEnviroment(appEnvironment)

	// Validate database connection
	checkParameterConnection(databaseConnection)

	// Open connection and Do Migrate
	dbConnect := openDatabaseConnection(databaseConnection)
	errUserMigration := dbConnect.AutoMigrate(&Migrations.User{})
	if errUserMigration != nil {
		log.Panic("Error when do migration, please refer this stack trace to debugging: " + errUserMigration.Error())
	}
}

// checkAppEnviroment will check APP_ENV constant in .env file
func checkAppEnviroment(appEnvironment string) {
	if appEnvironment == "" {
		log.Panic("APP_ENV cannot be empty")
	} else {
		fmt.Printf("Migration running in %v environment \n", appEnvironment)
	}
}

// checkParameterConnection will check DB Connection param from .env
// to make sure primary parameter not have an empty value
func checkParameterConnection(parameterConnection *DatabaseConnection) {
	if parameterConnection.DBHost == "" {
		log.Panic("DB_HOST cannot be empty")
	}

	if parameterConnection.DBUsername == "" {
		log.Panic("DB_USERNAME cannot be empty")
	}

	if parameterConnection.DBPassword == "" {
		log.Panic("DB_PASSWORD cannot be empty")
	}

	if parameterConnection.DBSchema == "" {
		log.Panic("DB_SCHEMA cannot be empty")
	}
}

// openDatabaseConnection used to create connection between command to database
func openDatabaseConnection(parameterConnection *DatabaseConnection) *gorm.DB {
	dsn := parameterConnection.DBUsername + ":" +
		parameterConnection.DBPassword + "@(" +
		parameterConnection.DBHost + ")/" +
		parameterConnection.DBSchema

	openConnection, errOpenConnection := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errOpenConnection != nil {
		log.Panic("Error when try to connect database, please refer this stack trace for debugging: \n",
			errOpenConnection.Error())
	}

	return openConnection
}
