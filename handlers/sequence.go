package handlers

import (
	"net/http"
	"strconv"

	"genomic-api/config"
	"genomic-api/models"

	"github.com/gin-gonic/gin"
)

// ListSequenceFiles godoc
// @Summary      List sequence files
// @Description  Get all sequence files
// @Tags         sequence
// @Produce      json
// @Success      200  {array}  models.SequenceFile
// @Router       /api/sequence [get]
func ListSequenceFiles(c *gin.Context) {
	var files []models.SequenceFile
	if err := config.DB.Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

// CreateSequenceFile godoc
// @Summary      Create sequence file
// @Description  Add a new sequence file record
// @Tags         sequence
// @Accept       json
// @Produce      json
// @Param        sequence_file  body  models.SequenceFile  true  "Sequence file info"
// @Success      201  {object}  models.SequenceFile
// @Router       /api/sequence [post]
func CreateSequenceFile(c *gin.Context) {
	var file models.SequenceFile
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, file)
}

// GetSequenceFile godoc
// @Summary      Get sequence file
// @Description  Get sequence file by ID
// @Tags         sequence
// @Produce      json
// @Param        id   path      int  true  "Sequence file ID"
// @Success      200  {object}  models.SequenceFile
// @Failure      404  {object}  map[string]string
// @Router       /api/sequence/{id} [get]
func GetSequenceFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var file models.SequenceFile
	if err := config.DB.First(&file, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sequence file not found"})
		return
	}
	c.JSON(http.StatusOK, file)
}

// UpdateSequenceFile godoc
// @Summary      Update sequence file
// @Description  Update sequence file by ID
// @Tags         sequence
// @Accept       json
// @Produce      json
// @Param        id             path      int                true  "Sequence file ID"
// @Param        sequence_file  body      models.SequenceFile true  "Sequence file info"
// @Success      200            {object}  models.SequenceFile
// @Failure      404            {object}  map[string]string
// @Router       /api/sequence/{id} [put]
func UpdateSequenceFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var file models.SequenceFile
	if err := config.DB.First(&file, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sequence file not found"})
		return
	}
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, file)
}

// DeleteSequenceFile godoc
// @Summary      Delete sequence file
// @Description  Delete sequence file by ID
// @Tags         sequence
// @Produce      json
// @Param        id   path      int  true  "Sequence file ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/sequence/{id} [delete]
func DeleteSequenceFile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.SequenceFile{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sequence file deleted"})
}
