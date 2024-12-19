package main

import (
	"log"
	"museum-api/controllers"
	"museum-api/database"
	"museum-api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()
}

func main() {
	initConfig()
	database.InitDB()
	database.RunMigrations()

	// Inicia o router Gin
	r := gin.Default()

	// Adiciona o middleware de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Permitir o domínio do frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Endpoints públicos (sem autenticação)
	r.POST("/api/v1/managers", controllers.CreateManager)
	r.POST("/api/v1/login", controllers.Login)
	r.GET("/api/v1/museums/city/:city", controllers.GetMuseumsByCity)
	r.GET("/api/v1/museums/state/:state", controllers.GetMuseumsByState)
	r.GET("/api/v1/museums/name/:name", controllers.GetMuseumsByName)
	r.GET("api/v1/artworks/museum/:name", controllers.GetArtworksByMuseumName)
	r.GET("api/v1/artworks/author/:author", controllers.GetArtworksByAuthor)
	r.GET("api/v1/artworks/name/:name", controllers.GetArtworksByName)
	r.GET("api/v1/artworks/year/:year", controllers.GetArtworksByYear)
	r.GET("/api/v1/museums", controllers.GetAllMuseums)
	r.GET("/api/v1/museums/category", controllers.GetMuseumsByCategory)

	// Endpoints protegidos por autenticação (usando middleware de validação de token)
	auth := r.Group("/api/v1")
	auth.Use(utils.ValidateTokenMiddleware)

	auth.POST("/museums", controllers.CreateMuseum)
	auth.POST("/artworks", controllers.CreateArtwork)
	auth.PUT("/museums/:id", controllers.UpdateMuseum)
	auth.PUT("/managers/:id", controllers.UpdateManager)
	auth.PUT("/museums/:id/disable", controllers.DisableMuseum)
	auth.PUT("/managers/:id/disable", controllers.DisableManager)
	auth.PUT("/artworks/:id/disable", controllers.DisableArtwork)
	auth.PUT("/artworks/:id", controllers.UpdateArtwork)

	// Inicializa o servidor na porta 3000
	r.Run(":3000")
}
