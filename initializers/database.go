package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectToDB() {
	dsn := os.Getenv("DB_CONNECTION")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

	fmt.Print("connected", DB)

	

}
