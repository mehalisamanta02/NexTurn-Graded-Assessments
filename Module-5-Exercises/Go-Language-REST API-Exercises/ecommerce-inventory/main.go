package main

import (
	"ecommerce-inventory/controllers"
	"ecommerce-inventory/dbconfig"
	"ecommerce-inventory/middlewares"
	"ecommerce-inventory/repositories"
	"ecommerce-inventory/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := dbconfig.InitializeDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set up repositories, services, and controllers
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Set up router
	router := gin.Default()

	// Middleware for logging requests
	router.Use(middlewares.LoggingMiddleware())

	// User routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Product routes (authentication required)
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware()) // Middleware for authentication
	{
		// Routes for managing products
		authorized.POST("/product", middlewares.ValidationMiddleware(), productController.AddProduct)
		authorized.GET("/product/:id", productController.GetProduct)
		authorized.PUT("/product/:id", productController.UpdateProduct)
		authorized.DELETE("/product/:id", productController.DeleteProduct)
		authorized.GET("/products", productController.GetAllProducts)
	}

	// Start the server on port 8080
	router.Run(":8080")
}
