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
// @Summary Factorización QR de una matriz
// @Description Calcula las matrices Q y R para una matriz dada y guarda el historial.
// @Tags matrix
// @Accept json
// @Produce json
// @Param matrix body dtos.QRRequest true "Matriz a factorizar"
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

	// Obtener el userID del contexto (inyectado por el middleware)
	userID := c.Locals("user_id").(string)

	result, err := ctrl.service.FactorizeQR(userID, req.Matrix)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(dtos.NewSuccessResponse(fiber.StatusOK, "Factorización QR completada", result))
}
