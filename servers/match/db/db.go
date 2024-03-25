package db

import (
	"log"
	"match/schema"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	AutoMigrate(models ...interface{}) error
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
}

// Ensure that *gorm.DB satisfies the Database interface
var _ Database = (*gorm.DB)(nil)

func InitDB() *gorm.DB {
	DB := ConnectPostgresDB()
	AutoMigratePostgresDB(DB)

	return DB
}

// ConnectPostgresDB initializes and returns a GORM database connection
func ConnectPostgresDB() *gorm.DB {
	// Load the database URL from the environment
	dburl := os.Getenv("DATABASE_URL")

	// Open the connection
	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
		return nil
	}

	log.Println("Successfully connected to PostgreSQL")
	return db
}

// AutoMigratePostgresDB migrates the database schema
func AutoMigratePostgresDB(db *gorm.DB) error {
	// Migrate the schema
	err := db.AutoMigrate(&schema.Match{})
	if err != nil {
		log.Fatalf("Error migrating PostgreSQL schema: %v", err)
		return err
	}

	log.Println("Successfully migrated PostgreSQL schema")
	return nil
}
