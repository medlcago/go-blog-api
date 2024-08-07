package postgres

import (
	"fmt"
	"go-blog-api/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func ConnectDB(host, port, user, name, password string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, name, password)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	return db
}
