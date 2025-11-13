package handlers

import (
	"database/sql"
	"errors"
	"personcrud/internal/models"
	"personcrud/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	service *services.PersonService
}

func NewPersonHandler(service *services.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) GetAll(c *gin.Context) {
	persons, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, persons)
}

func (h *PersonHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	person, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, services.ErrPersonNotFound) {
			c.JSON(404, gin.H{"error": "Person not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if person == nil {
		c.JSON(404, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(200, person)
}

func (h *PersonHandler) Create(c *gin.Context) {
	var req models.CreatePersonRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	id, err := h.service.Create(&req)
	if err != nil {
		// Проверяем тип ошибки для возврата соответствующего статуса
		if errors.Is(err, services.ErrInvalidEmail) ||
			errors.Is(err, services.ErrInvalidPhone) ||
			errors.Is(err, services.ErrInvalidFirstName) ||
			errors.Is(err, services.ErrInvalidLastName) {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to create person"})
		return
	}

	c.JSON(201, gin.H{
		"id":      id,
		"message": "Person created successfully",
	})
}

func (h *PersonHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var req models.CreatePersonRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	err = h.service.Update(id, &req)
	if err != nil {
		// Проверяем тип ошибки для возврата соответствующего статуса
		if errors.Is(err, services.ErrInvalidEmail) ||
			errors.Is(err, services.ErrInvalidPhone) ||
			errors.Is(err, services.ErrInvalidFirstName) ||
			errors.Is(err, services.ErrInvalidLastName) {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, services.ErrPersonNotFound) {
			c.JSON(404, gin.H{"error": "Person not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to update person"})
		return
	}

	c.JSON(200, gin.H{"message": "Person updated successfully"})
}

func (h *PersonHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, services.ErrPersonNotFound) {
			c.JSON(404, gin.H{"error": "Person not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to delete person"})
		return
	}

	c.JSON(200, gin.H{"message": "Person deleted successfully"})
}
