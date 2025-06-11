package controllers

import (
	"net/http"
	"strconv"

	"github.com/SidharthaDR/golang-clinic-app/internal/models"
	"github.com/SidharthaDR/golang-clinic-app/internal/utils"
	"github.com/gin-gonic/gin"
)

// CreatePatient registers a new patient
func CreatePatient(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//inp validation backend
	if patient.Name == "" || patient.Age <= 0 || len(patient.Phone) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient data"})
		return
	}

	if err := utils.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

// GetPatients returns all patients
func GetPatients(c *gin.Context) {
	var patients []models.Patient
	if err := utils.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

// GetPatient returns a single patient by ID
func GetPatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var patient models.Patient
	if err := utils.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient updates patient data
func UpdatePatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var patient models.Patient

	if err := utils.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if patient.Name == "" || patient.Age <= 0 || len(patient.Phone) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient data"})
		return
	}

	if err := utils.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// DeletePatient deletes a patient by ID
func DeletePatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := utils.DB.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
