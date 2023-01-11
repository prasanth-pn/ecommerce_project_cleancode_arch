package db

import (
	config "clean/pkg/config"
	domain "clean/pkg/domain"

	//"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGormDB(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo:="host=localhost user=tuser dbname=sample port=5432 password=1234"
	psqlInfo := cfg.DBSOURCE
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(domain.Admins{})
	db.AutoMigrate(domain.Product{})
	db.AutoMigrate(domain.Category{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Teacher{})
	db.AutoMigrate(&domain.Student{})
	return db, dbErr
}
