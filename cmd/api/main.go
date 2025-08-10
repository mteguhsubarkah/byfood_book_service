package main

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "byfood_service/internal/config"
    "byfood_service/internal/domain"
    "byfood_service/internal/service"
    "byfood_service/internal/handler"
    "byfood_service/internal/http"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

func main() {
    // Load config
    err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Build DSN 
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        config.Cfg.GetString(config.DbHost),
        config.Cfg.GetString(config.DbUser),
        config.Cfg.GetString(config.DbPassword),
        config.Cfg.GetString(config.DbName),
        config.Cfg.GetString(config.DbPort),
        config.Cfg.GetString(config.DbSSLMode),
        config.Cfg.GetString(config.DbTimezone),
    )

    // Connecting and migrating database
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err := db.AutoMigrate(&domain.Book{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    r := gin.Default()

    // Setup CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    // Logging middleware
    r.Use(func(c *gin.Context) {
        log.Printf("Incoming %s %s from %s", c.Request.Method, c.Request.RequestURI, c.ClientIP())
        c.Next()
        log.Printf("Completed %s %s with status %d", c.Request.Method, c.Request.RequestURI, c.Writer.Status())
    })

	// Iniitialize service with db
	service := service.NewBookService(db)

	// Initialize handler with service
	handler := handler.NewBookHandler(service)

	// Register API routes
	http.Route(r, handler)


    // Load service port and start the service
    port := config.Cfg.GetString(config.ServicePort)
    if port == "" {
        port = "8080"
    }

    log.Printf("byfood_service running at http://localhost:%s", port)
    r.Run(fmt.Sprintf(":%s", port))
}
