package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rquiogue/travel-to-do-list/internal/models"
	"github.com/rquiogue/travel-to-do-list/internal/repositories"
)

type LocationController struct {
	Repo *repositories.LocationRepository
}

func NewLocationController(repo *repositories.LocationRepository) *LocationController {
	return &LocationController{Repo: repo}
}

func (c *LocationController) GetLocations(ctx *gin.Context) {
	locations, err := c.Repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, locations)
}

func (c *LocationController) CreateLocation(ctx *gin.Context) {
	var location models.Location
	if err := ctx.BindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if location.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	location.Completed = false

	if err := c.Repo.Create(&location); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, location)
}

func (c *LocationController) UpdateLocation(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var location models.Location
	if err := ctx.BindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	location.ID = id

	if err := c.Repo.Update(&location); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, location)
}

func (c *LocationController) DeleteLocation(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.Repo.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
