package store

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)

var DB *gorm.DB

// Setup initializes the database instance
func DBSetup() {
	var err error
	// Update the connection string format for PostgreSQL
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Port,
		config.DatabaseSetting.User,
		config.DatabaseSetting.Name,
		config.DatabaseSetting.Password))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DatabaseSetting.TablePrefix + defaultTableName
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

    log.Println("DB Connected")



}

func CloseDBConn() {
	if DB == nil {
		return
	}
    err := DB.Close()
    if err != nil{
        fmt.Printf("failed to close database connection")
    }
}
