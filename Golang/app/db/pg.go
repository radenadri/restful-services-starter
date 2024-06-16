package db

import (
	"boilerplate/app/config"
	"boilerplate/app/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_SSL_MODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("Failed to connect database")
	}

	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

func Paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
