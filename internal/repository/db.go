package repository

import (
	"final-project-1/internal/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitDB(host, port, user, password, dbname, dialect string) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	

	db, err := gorm.Open(dialect, dbURI)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	// Migrate database schema
	err = db.AutoMigrate(&domain.Todo{}).Error
	if err != nil {
		return nil, err
	}

	return db, nil
}
