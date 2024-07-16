package main

import (
	"log"
	"net/http"
	"github.com/yuvakkrishnan/user-service/internal/handlers"
	"github.com/yuvakkrishnan/user-service/internal/services"
	"github.com/yuvakkrishnan/user-service/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
	log.Printf("Entering into main func")
	 // Replace with your actual PostgreSQL connection string
	 dsn := "host=localhost user=postgres password=postgres dbname=userservice port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	 if err != nil {
		 panic("failed to connect database")
	 }
 
	 // Migrate the schema
	 db.AutoMigrate(&models.User{})
 
	 // Initialize user service with the database connection
	 userService := services.NewUserService(db)
 
	 // Start your application with the userService and HTTP server setup
	 http.ListenAndServe(":8080", handlers.NewRouter(userService))

	log.Println("Starting server on :8080")
}