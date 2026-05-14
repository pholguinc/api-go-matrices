package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pholguinc/api-go-matrices/internal/constants"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
	"github.com/pholguinc/api-go-matrices/internal/services"
)

type MatrixController struct {
	service services.MatrixService
}

func NewMatrixController(service services.MatrixService) *MatrixController {
	return &MatrixController{service: service}
}

// Factorize godoc
// @Summary Realiza la factorización QR
// @Description Calcula las matrices Q (ortogonal) y R (triangular superior) de una matriz dada.
// @Tags matrix
// @Accept json
// @Produce json
// @Param matrix body dtos.QRRequest true "Matriz original"
// @Success 200 {object} dtos.ApiResponse{data=dtos.QRResponseData}
// @Failure 400 {object} dtos.ApiResponse
// @Failure 500 {object} dtos.ApiResponse
// @Security BearerAuth
// @Router /matrix/factorize [post]
func (ctrl *MatrixController) Factorize(c fiber.Ctx) error {
	var req dtos.QRRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.NewErrorResponse(fiber.StatusBadRequest, constants.ErrInvalidRequestBody))
	}

	if len(req.Matrix) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.NewErrorResponse(fiber.StatusBadRequest, constants.ErrMatrixEmpty))
	}

	result, err := ctrl.service.FactorizeQR(req.Matrix)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.NewErrorResponse(fiber.StatusInternalServerError, "Error en el cálculo: "+err.Error()))
	}

	return c.JSON(dtos.NewSuccessResponse(fiber.StatusOK, "Factorización QR completada", result))
}
