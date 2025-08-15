package http

import (
	"finance/internal/usecase/installment"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstallmentHandler struct {
	uc installment.UseCase
}

func NewInstallmentHandler(uc installment.UseCase) *InstallmentHandler {
	return &InstallmentHandler{uc: uc}
}

func (h *InstallmentHandler) CreateInstallment(c *gin.Context) {
	var input installment.CreateInstallmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdInstallment, err := h.uc.CreateInstallment(c.Request.Context(), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdInstallment)
}
