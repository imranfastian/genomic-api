package handlers

import (
	"net/http"
	"strconv"

	"genomic-api/config"
	"genomic-api/models"

	"github.com/gin-gonic/gin"
)

// ListVariants godoc
// @Summary      List variant files
// @Description  Get all variant files
// @Tags         variants
// @Produce      json
// @Success      200  {array}  models.VariantFile
// @Router       /api/variants [get]
func ListVariants(c *gin.Context) {
	var variants []models.VariantFile
	if err := config.DB.Find(&variants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variants)
}

// CreateVariant godoc
// @Summary      Create variant file
// @Description  Add a new variant file record
// @Tags         variants
// @Accept       json
// @Produce      json
// @Param        variant_file  body  models.VariantFile  true  "Variant file info"
// @Success      201  {object}  models.VariantFile
// @Router       /api/variants [post]
func CreateVariant(c *gin.Context) {
	var variant models.VariantFile
	if err := c.ShouldBindJSON(&variant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&variant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, variant)
}

// GetSampleVariants godoc
// @Summary      Get sample variants
// @Description  Get all variant files for a sample
// @Tags         variants
// @Produce      json
// @Param        id   path      int  true  "Sample ID"
// @Success      200  {array}  models.VariantFile
// @Router       /api/samples/{id}/variants [get]
func GetSampleVariants(c *gin.Context) {
	sampleID, _ := strconv.Atoi(c.Param("id"))
	var variants []models.VariantFile
	if err := config.DB.Where("sample_id = ?", sampleID).Find(&variants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variants)
}

// DeleteVariant godoc
// @Summary      Delete variant file
// @Description  Delete variant file by ID
// @Tags         variants
// @Produce      json
// @Param        id   path      int  true  "Variant file ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/variants/{id} [delete]
func DeleteVariant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.VariantFile{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant deleted"})
}
