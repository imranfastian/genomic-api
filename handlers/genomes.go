package handlers

import (
	"net/http"
	"strconv"

	"genomic-api/config"
	"genomic-api/models"

	"github.com/gin-gonic/gin"
)

// ListGenomes godoc
// @Summary      List genomes
// @Description  Get all genomes
// @Tags         genomes
// @Produce      json
// @Success      200  {array}  models.Genome
// @Router       /api/genomes [get]
func ListGenomes(c *gin.Context) {
	var genomes []models.Genome
	if err := config.DB.Find(&genomes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genomes)
}

// CreateGenome godoc
// @Summary      Create genome
// @Description  Add a new genome record
// @Tags         genomes
// @Accept       json
// @Produce      json
// @Param        genome  body  models.Genome  true  "Genome info"
// @Success      201  {object}  models.Genome
// @Router       /api/genomes [post]
func CreateGenome(c *gin.Context) {
	var genome models.Genome
	if err := c.ShouldBindJSON(&genome); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&genome).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, genome)
}

// GetGenome godoc
// @Summary      Get genome
// @Description  Get genome by ID
// @Tags         genomes
// @Produce      json
// @Param        id   path      int  true  "Genome ID"
// @Success      200  {object}  models.Genome
// @Failure      404  {object}  map[string]string
// @Router       /api/genomes/{id} [get]
func GetGenome(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var genome models.Genome
	if err := config.DB.First(&genome, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genome not found"})
		return
	}
	c.JSON(http.StatusOK, genome)
}

// UpdateGenome godoc
// @Summary      Update genome
// @Description  Update genome by ID
// @Tags         genomes
// @Accept       json
// @Produce      json
// @Param        id      path      int           true  "Genome ID"
// @Param        genome  body      models.Genome true  "Genome info"
// @Success      200     {object}  models.Genome
// @Failure      404     {object}  map[string]string
// @Router       /api/genomes/{id} [put]
func UpdateGenome(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var genome models.Genome
	if err := config.DB.First(&genome, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genome not found"})
		return
	}
	if err := c.ShouldBindJSON(&genome); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&genome).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genome)
}

// DeleteGenome godoc
// @Summary      Delete genome
// @Description  Delete genome by ID
// @Tags         genomes
// @Produce      json
// @Param        id   path      int  true  "Genome ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/genomes/{id} [delete]
func DeleteGenome(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Genome{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Genome deleted"})
}
