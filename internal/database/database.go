// package database

// import (
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var DBConn *gorm.DB

// func ConnectDB () {
// 	dsn := "postgresql://postgres:16101227@bracu@db.tiocojqjctxlckestsmr.supabase.co:5432/postgres"
//   	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Error),
// 	})

// 	if err != nil {
// 		panic("Database connection failed!")
// 	}
// 	log.Println("Connection Success!")

// 	DBConn = db

// }

package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global connection variable other packages (like routes or controllers) will use
var DB *gorm.DB

func ConnectDB() {
	// 1. Define your local connection credentials
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "01676022659" 
	dbname := "postgres"         

	// 2. Build the GORM-compatible Data Source Name (DSN) string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka", 
		host, user, password, dbname, port)

	// 3. Open the GORM connection (Notice we DO NOT defer close here)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Shows SQL queries in terminal
	})

	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Database connection established successfully!")

	// 4. Assign the local connection instance to our global variable
	DB = db
}
