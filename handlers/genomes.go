
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "genomic-api/config"
    "genomic-api/models"
)

func ListGenomes(c *gin.Context) {
    var genomes []models.Genome
    if err := config.DB.Find(&genomes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, genomes)
}

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

func GetGenome(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var genome models.Genome
    if err := config.DB.First(&genome, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Genome not found"})
        return
    }
    c.JSON(http.StatusOK, genome)
}

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

func DeleteGenome(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := config.DB.Delete(&models.Genome{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Genome deleted"})
}
