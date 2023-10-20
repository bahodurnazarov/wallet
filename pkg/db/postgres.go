package db

import (
	"database/sql"
	"fmt"
	"github.com/bahodurnazarov/wallet/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func ConnectDB(cfg *models.Settings) (*gorm.DB, error) {
	db, err := postgresDB(cfg)
	return db, err

}
func postgresDB(cfg *models.Settings) (*gorm.DB, error) {
	database := cfg.Postgres

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database.PG_host, database.PG_port, database.PG_user, database.PG_password, database.PG_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.User{}, &models.Transactions{}); err != nil {
		log.Println(err.Error())
	}

	DB = db

	return DB, nil
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// Close
	return sqlDB.Close()
}
