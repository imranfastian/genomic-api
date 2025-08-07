package handlers

import (
	"net/http"
	"strconv"

	"genomic-api/config"
	"genomic-api/models"

	"github.com/gin-gonic/gin"
)

// ListSamples godoc
// @Summary      List samples
// @Description  Get all samples
// @Tags         samples
// @Produce      json
// @Success      200  {array}  models.Sample
// @Router       /api/samples [get]
func ListSamples(c *gin.Context) {
	var samples []models.Sample
	if err := config.DB.Find(&samples).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, samples)
}

// CreateSample godoc
// @Summary      Create sample
// @Description  Add a new sample record
// @Tags         samples
// @Accept       json
// @Produce      json
// @Param        sample  body  models.Sample  true  "Sample info"
// @Success      201  {object}  models.Sample
// @Router       /api/samples [post]
func CreateSample(c *gin.Context) {
	var sample models.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&sample).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sample)
}

// GetSample godoc
// @Summary      Get sample
// @Description  Get sample by ID
// @Tags         samples
// @Produce      json
// @Param        id   path      int  true  "Sample ID"
// @Success      200  {object}  models.Sample
// @Failure      404  {object}  map[string]string
// @Router       /api/samples/{id} [get]
func GetSample(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sample models.Sample
	if err := config.DB.First(&sample, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sample not found"})
		return
	}
	c.JSON(http.StatusOK, sample)
}

// UpdateSample godoc
// @Summary      Update sample
// @Description  Update sample by ID
// @Tags         samples
// @Accept       json
// @Produce      json
// @Param        id      path      int           true  "Sample ID"
// @Param        sample  body      models.Sample true  "Sample info"
// @Success      200     {object}  models.Sample
// @Failure      404     {object}  map[string]string
// @Router       /api/samples/{id} [put]
func UpdateSample(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sample models.Sample
	if err := config.DB.First(&sample, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sample not found"})
		return
	}
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&sample).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sample)
}

// DeleteSample godoc
// @Summary      Delete sample
// @Description  Delete sample by ID
// @Tags         samples
// @Produce      json
// @Param        id   path      int  true  "Sample ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/samples/{id} [delete]
func DeleteSample(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Sample{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sample deleted"})
}
