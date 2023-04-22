package primedb

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	client *gorm.DB
}

type Settings struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func Connect(s Settings) (DB, error) {

	client, err := gorm.Open(postgres.Open(
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=UTC",
			s.User,
			s.Password,
			s.Database,
			s.Host,
			s.Port)), &gorm.Config{})
	if err != nil {
		return DB{}, err
	}

	return DB{
		client: client,
	}, nil
}

func (db *DB) Close() error {
	c, _ := db.client.DB()
	return c.Close()
}
