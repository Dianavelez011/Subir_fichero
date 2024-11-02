package config

import (
    "log"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func SetupDB() *gorm.DB {
    dsn := "host=localhost user=admin password=admin123 dbname=gestion_giras port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    return db
}