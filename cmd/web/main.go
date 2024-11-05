package main

import (
	authhandler "bookingcinema/pkg/auth/handlers"
	authinfra "bookingcinema/pkg/auth/infrastructe/postgres"
	authusecase "bookingcinema/pkg/auth/usecases"
	authutils "bookingcinema/pkg/auth/utils"
	dbpg "bookingcinema/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load .env file")
	}

	conn, err := dbpg.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer conn.Close()

	// Initialize Fiber
	app := fiber.New()

	// Repositories and Services
	userRepo := authinfra.NewUserRepository(conn)
	authUsecase := authusecase.NewAuthenticationUseCase(userRepo)
	authHandler := authhandler.NewAuthHandler(authUsecase)

	api := app.Group("/api")

	v1 := api.Group("v1")

	// Auth Routes
	v1.Post("/auth/register", authHandler.RegisterHandler)
	v1.Post("/auth/login", authHandler.LoginHandler)
	v1.Get("/auth/user", authutils.AuthMiddleware(userRepo), authHandler.UserHandler)

	// User can view their own roles and permissions
	//v1.Get("/user/roles", authutils.AuthMiddleware(userRepo), userHandler.ViewOwnRoles)
	//v1.Get("/user/permissions", authutils.AuthMiddleware(userRepo), userHandler.ViewOwnPermissions)

	// Admin-only routes
	//app.Post("/api/admin/assign-role", authutils.AuthMiddleware(userRepo), authutils.AdminOnlyMiddleware(), adminHandler.AssignRoleToUser)
	//app.Get("/api/admin/user/:id/roles", authutils.AuthMiddleware(userRepo), authutils.AdminOnlyMiddleware(), adminHandler.ViewUserRoles)

	// Start the Fiber server
	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
