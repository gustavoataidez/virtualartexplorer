package handler

import (
	"log"
	"museum-api/api/controllers"
	"museum-api/api/database"
	"museum-api/api/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		// Se o .env não for encontrado, apenas avise (não é crítico)
		println("No .env file found.")
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
}

func setupRouter() *gin.Engine {
	initConfig()
	database.InitDB()
	database.RunMigrations()

	r := gin.Default()

	// Configuração de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Permitir o domínio do frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
	
	// Endpoints públicos (sem autenticação)
	r.POST("/api/v1/managers", controllers.CreateManager)
	r.POST("/api/v1/login", controllers.Login)
	r.GET("/api/v1/museums/city/:city", controllers.GetMuseumsByCity)
	r.GET("/api/v1/museums/state/:state", controllers.GetMuseumsByState)
	r.GET("/api/v1/museums/name/:name", controllers.GetMuseumsByName)
	r.GET("/api/v1/artworks/museum/name/:name", controllers.GetArtworksByMuseumName)
	r.GET("/api/v1/artworks/author/:author", controllers.GetArtworksByAuthor)
	r.GET("/api/v1/artworks/name/:name", controllers.GetArtworksByName)
	r.GET("/api/v1/artworks/year/:year", controllers.GetArtworksByYear)
	r.GET("/api/v1/museums", controllers.GetAllMuseums)
	r.GET("/api/v1/museums/category", controllers.GetMuseumsByCategory)
	r.GET("/api/v1/museums/:id", controllers.GetMuseumByID)
	r.GET("/api/v1/artworks", controllers.GetAllArtworks)
	r.GET("/api/v1/artworks/museum/id/:id", controllers.GetArtworkByMuseumId)
	r.GET("/api/v1/museums/category/:category", controllers.GetMuseumsByCategory1)

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

	// Adiciona a nova rota para listar museus do usuário autenticado
	auth.GET("/museums/my", controllers.GetMuseumsByAuthenticatedUser)

	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received request: %s %s", r.Method, r.URL.Path)
    router := setupRouter()
    router.ServeHTTP(w, r)
}
