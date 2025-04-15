package db

import (
    "log"

    "github.com/rmrachmanfauzan/elpempek/internal/model"
)

// RunMigrations migrates the database schemas
func RunMigrations() {
    err := DB.AutoMigrate(
        &model.User{},
    )
    if err != nil {
        log.Fatalf("❌ Migration failed: %v", err)
    }

    log.Println("✅ Database migrated successfully")
}
