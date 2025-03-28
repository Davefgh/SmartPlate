package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"vehicle-api/internal/database"
	"vehicle-api/internal/handlers"
	"vehicle-api/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	rand.Seed(time.Now().UnixNano())
	// Initialize repositories and handlers
	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)


	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Enhanced CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           3600,
	}))
	// Vehicle routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})
	e.GET("/generate-lto-id", userHandler.GenerateLTOID)
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.GET("/users/email/:email", userHandler.GetUserByEmail)
	e.PUT("/users/:id", userHandler.UpdateUser)	
	//for getting user by lto client id
	e.GET("/users/lto/:lto_client_id", userHandler.GetUserByLTOClientID)
	e.PUT("/users/by-lto/:lto_client_id", userHandler.UpdateUserByLTO)
	e.DELETE("/users/by-lto/:lto_client_id", userHandler.DeleteUserByLTO)



	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}