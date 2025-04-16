package db

import (
	"os"
	"fmt"
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDB() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		// fallback if env is not set
		// dsn = "root:superadmin@tcp(db:3306)/elpempek?charset=utf8mb4&parseTime=True&loc=Local"
		dsn = "root:superadmin@tcp(127.0.0.1:3306)/elpempek?charset=utf8mb4&parseTime=True&loc=Local"
	}
	var db *gorm.DB	
	var err error
	for i := 0; i < 10; i++ {
		db , err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			fmt.Println("✅ Database connection established")
			return
		}
		log.Printf("Retrying database connection (%d/10): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("❌ Failed to connect to database after retries: %v", err)
}
