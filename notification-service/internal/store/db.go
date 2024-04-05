package store

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL Gorm dialect

    "github.com/dranonymousnet/eventhorizon/internal/config"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
    var err error
    // Update the connection string format for PostgreSQL
    db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
        config.DatabaseSetting.Host,
        config.DatabaseSetting.User,
        config.DatabaseSetting.Name,
        config.DatabaseSetting.Password))

    if err != nil {
        log.Fatalf("models.Setup err: %v", err)
    }

    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return config.DatabaseSetting.TablePrefix + defaultTableName
    }
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}
