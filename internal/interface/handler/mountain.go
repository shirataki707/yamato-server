package handler

import (
	"net/http"
	"strconv"
	"yamato/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MountainHandler struct {
	useCase usecase.MountainUseCase
}

func NewMountainHandler(useCase usecase.MountainUseCase) *MountainHandler {
	return &MountainHandler{useCase: useCase}
}

func (h *MountainHandler) GetMountain(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mountain id"})
		return
	}

	mountain, err := h.useCase.FetchMountainByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mountain)
}

func (h *MountainHandler) GetAllMountains(c *gin.Context) {
	mountains, err := h.useCase.FetchAllMountains()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mountains)
}
