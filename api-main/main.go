package main

import (
	"log"
	"museum-api/controllers"
	"museum-api/database"
	"museum-api/utils"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Lê DB_URL e PORT das variáveis de ambiente
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // valor padrão
	}

	// Inicializa DB e migrações
	database.InitDB()
	database.RunMigrations()

	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://visualartexplorer.vercel.app", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rotas públicas
	r.POST("/api/v1/managers", controllers.CreateManager)
	r.POST("/api/v1/login", controllers.Login)
	r.GET("/api/v1/museums/city/:city", controllers.GetMuseumsByCity)
	r.GET("/api/v1/museums/state/:state", controllers.GetMuseumsByState)
	r.GET("/api/v1/museums/name/:name", controllers.GetMuseumsByName)
	r.GET("/api/v1/artworks/museum/:name", controllers.GetArtworksByMuseumName)
	r.GET("/api/v1/artworks/author/:author", controllers.GetArtworksByAuthor)
	r.GET("/api/v1/artworks/name/:name", controllers.GetArtworksByName)
	r.GET("/api/v1/artworks/year/:year", controllers.GetArtworksByYear)
	r.GET("/api/v1/museums", controllers.GetAllMuseums)
	r.GET("/api/v1/museums/category", controllers.GetMuseumsByCategory)
	r.GET("/api/v1/museums/:id", controllers.GetMuseumByID)
	r.GET("/api/v1/artworks", controllers.GetAllArtworks)
	r.GET("/api/v1/artworks/museum/id/:id", controllers.GetArtworkByMuseumId)
	r.GET("/api/v1/museums/category/:category", controllers.GetMuseumsByCategory1)

	// Rotas autenticadas
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
	auth.GET("/museums/my", controllers.GetMuseumsByAuthenticatedUser)

	// 🚀 Porta configurável
	r.Run(":" + port)
}
