package controllers

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"museum-api/database"
	"museum-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllArtworks(c *gin.Context) {
	// Variável para armazenar todas as obras de arte
	var artworks []models.Artwork

	// Busca todas as obras de arte no banco de dados
	if err := database.DB.Find(&artworks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch artworks"})
		return
	}

	// Retorna as obras de arte
	c.JSON(http.StatusOK, artworks)
}


func CreateArtwork(c *gin.Context) {
	var artwork models.Artwork

	// Obtém o título (obrigatório)
	artwork.Name = c.PostForm("name")
	if artwork.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Obtém o museum_id (opcional)
	museumIDStr := c.PostForm("museum_id")
	if museumIDStr != "" {
		museumID, err := strconv.Atoi(museumIDStr)
		if err != nil || museumID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid museum_id"})
			return
		}
		artwork.MuseumID = uint(museumID)

		// Verifica se o museu existe
		var museum models.Museum
		if err := database.DB.First(&museum, artwork.MuseumID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Museum not found"})
			return
		}
	}

	// Obtém a descrição (opcional)
	artwork.Description = c.PostForm("description")

	// Obtém o autor (opcional)
	artwork.Author = c.PostForm("author")

	// Obtém a imagem (opcional)
	file, err := c.FormFile("image")
	if err == nil {
		imageFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
			return
		}
		defer func(imageFile multipart.File) {
			_ = imageFile.Close()
		}(imageFile)

		imageBytes, err := ioutil.ReadAll(imageFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
			return
		}
		artwork.Image = imageBytes
	}

	// Salva a obra de arte no banco de dados
	if err := database.DB.Create(&artwork).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artwork created successfully", "artwork": artwork})
}


func UpdateArtwork(c *gin.Context) {
	var artwork models.Artwork
	if err := c.ShouldBind(&artwork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	artworkID := c.Param("id")

	// Handle image file
	file, _, err := c.Request.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}
	if file != nil {
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		imageBytes, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
			return
		}
		artwork.Image = imageBytes
	}

	if err := database.DB.Model(&models.Artwork{}).Where("id = ?", artworkID).Updates(artwork).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artwork updated successfully"})
}

// DisableArtwork disables an artwork, AINDA É NECESSARIO IMPLEMENTAR VERIFICAÇÃO DE ALTERAÇÃO APOS A DESATIVAÇÃO
// Ja era pra ter feito isso seu arrombado, deixa de preguiça e faz logo
func DisableArtwork(c *gin.Context) {
	artworkID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artwork ID"})
		return
	}

	var artwork models.Artwork
	if err := database.DB.First(&artwork, artworkID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artwork not found"})
		return
	}

	artwork.Active = false
	if err := database.DB.Save(&artwork).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artwork disabled successfully"})
}

func GetArtworksByMuseumName(c *gin.Context) {
	museumName := c.Param("name")

	var museum models.Museum
	if err := database.DB.Where("name = ?", museumName).First(&museum).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Museum not found"})
		return
	}

	var artworks []models.Artwork
	if err := database.DB.Where("museum_id = ?", museum.ID).Find(&artworks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"artworks": artworks})
}

func GetArtworksByAuthor(c *gin.Context) {
	artist := c.Query("artist")

	var artworks []models.Artwork
	if err := database.DB.Where("author = ?", artist).Find(&artworks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"artworks": artworks})
}

func GetArtworksByYear(c *gin.Context) {
	year := c.Query("year")

	var artworks []models.Artwork
	if err := database.DB.Where("year = ?", year).Find(&artworks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"artworks": artworks})
}

func GetArtworksByName(c *gin.Context) {
	name := c.Query("name")

	var artworks []models.Artwork
	if err := database.DB.Where("name LIKE ?", "%"+name+"%").Find(&artworks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"artworks": artworks})
}
