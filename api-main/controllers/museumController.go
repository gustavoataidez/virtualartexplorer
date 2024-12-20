package controllers

import (
	"io/ioutil"
	"museum-api/database"
	"museum-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMuseumsByCategory1(c *gin.Context) {
	// Obtém a categoria dos parâmetros da rota
	category := c.Param("category")

	// Verifica se a categoria foi fornecida
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category1 is required"})
		return
	}

	// Variável para armazenar os museus encontrados
	var museums []models.Museum

	// Busca os museus no banco de dados associados à category1
	if err := database.DB.Where("category1 = ?", category).Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch museums"})
		return
	}

	// Retorna os museus encontrados
	c.JSON(http.StatusOK, museums)
}

func GetMuseumByID(c *gin.Context) {
	// Obtém o ID do museu a partir dos parâmetros da URL
	idParam := c.Param("id")

	// Converte o ID para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid museum ID"})
		return
	}

	// Variável para armazenar os dados do museu
	var museum models.Museum

	// Busca o museu no banco de dados
	if err := database.DB.First(&museum, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Museum not found"})
		return
	}

	// Retorna os dados do museu
	c.JSON(http.StatusOK, museum)
}

// GetMuseumsByAuthenticatedUser retorna os museus associados ao gerente autenticado
func GetMuseumsByAuthenticatedUser(c *gin.Context) {
	// Obtém o manager_id do contexto, definido pelo middleware
	managerID, exists := c.Get("manager_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var museums []models.Museum

	// Busca os museus no banco de dados associados ao manager_id
	if err := database.DB.Where("manager_id = ?", managerID).Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch museums"})
		return
	}

	// Retorna os museus encontrados
	c.JSON(http.StatusOK, museums)
}

func CreateMuseum(c *gin.Context) {
	var museum models.Museum

	// Bind obrigatório para title e description
	museum.Title = c.PostForm("title")
	museum.Description = c.PostForm("description")

	// Verifica se title e description foram fornecidos
	if museum.Title == "" || museum.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and description are required"})
		return
	}

	// Bind opcional para os outros campos
	museum.Category1 = c.PostForm("category1")
	museum.Category2 = c.PostForm("category2")
	museum.Link = c.PostForm("link")
	museum.Address = c.PostForm("address")
	museum.CEP = c.PostForm("cep")
	museum.City = c.PostForm("city")
	museum.State = c.PostForm("state")
	museum.Information = c.PostForm("information")

	// Processa imagem, se fornecida
	file, err := c.FormFile("image")
	if err == nil { // Apenas tenta processar a imagem se ela for enviada
		imageFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
			return
		}
		defer imageFile.Close()

		imageBytes, err := ioutil.ReadAll(imageFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
			return
		}
		museum.Image = imageBytes
	}

	// Pega o ID do gerente pelo token
	managerID, exists := c.Get("manager_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Manager ID not found in token"})
		return
	}
	museum.ManagerID = managerID.(uint)

	// Salva museu no banco de dados
	if err := database.DB.Create(&museum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Museu cadastrado com sucesso", "museum": museum})
}

func UpdateMuseum(c *gin.Context) {
	var updateData models.Museum
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	museumID := c.Param("id")
	var museum models.Museum
	if err := database.DB.Where("id = ?", museumID).First(&museum).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Museum not found"})
		return
	}

	if updateData.Title != "" {
		museum.Title = updateData.Title
	}
	if updateData.Description != "" {
		museum.Description = updateData.Description
	}
	if len(updateData.Image) > 0 {
		museum.Image = updateData.Image
	}
	if updateData.Category1 != "" {
		museum.Category1 = updateData.Category1
	}
	if updateData.Category2 != "" {
		museum.Category2 = updateData.Category2
	}
	if updateData.Link != "" {
		museum.Link = updateData.Link
	}
	if updateData.Address != "" {
		museum.Address = updateData.Address
	}
	if updateData.CEP != "" {
		museum.CEP = updateData.CEP
	}
	if updateData.City != "" {
		museum.City = updateData.City
	}
	if updateData.State != "" {
		museum.State = updateData.State
	}
	if updateData.Information != "" {
		museum.Information = updateData.Information
	}

	if err := database.DB.Save(&museum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Museum updated successfully"})
}

func GetAllMuseums(c *gin.Context) {
    var museums []models.Museum
    if err := database.DB.Find(&museums).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func GetMuseumsByCategory(c *gin.Context) {
    category := c.Query("category")

    var museums []models.Museum
    // Busca museus onde a categoria informada apareça em Category1 ou Category2
    if err := database.DB.Where("category1 = ? OR category2 = ?", category, category).Find(&museums).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func GetMuseumsByState(c *gin.Context) {
	state := c.Query("state")

	var museums []models.Museum
	if err := database.DB.Where("state = ?", state).Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func GetMuseumsByCity(c *gin.Context) {
	city := c.Query("city")

	var museums []models.Museum
	if err := database.DB.Where("city = ?", city).Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func GetMuseumsByName(c *gin.Context) {
	name := c.Query("name")

	var museums []models.Museum
	if err := database.DB.Where("title LIKE ?", "%"+name+"%").Find(&museums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"museums": museums})
}

func DisableMuseum(c *gin.Context) {
	museumID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid museum ID"})
		return
	}

	var museum models.Museum
	if err := database.DB.First(&museum, museumID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Museum not found"})
		return
	}

	museum.Active = false
	if err := database.DB.Save(&museum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Museum disabled successfully"})
}
