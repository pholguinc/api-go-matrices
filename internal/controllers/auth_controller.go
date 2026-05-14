package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
	"github.com/pholguinc/api-go-matrices/internal/services"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

// Register godoc
// @Summary Registro de nuevo usuario
// @Description Crea un nuevo usuario en el sistema con contraseña encriptada.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dtos.RegisterRequest true "Datos de registro"
// @Success 201 {object} dtos.ApiResponse "Usuario creado exitosamente"
// @Failure 400 {object} dtos.ApiResponse "Datos de entrada inválidos"
// @Failure 500 {object} dtos.ApiResponse "Error interno del servidor o email duplicado"
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c fiber.Ctx) error {
	var req dtos.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.NewErrorResponse(fiber.StatusBadRequest, "Datos inválidos"))
	}

	_, err := ctrl.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(dtos.NewSuccessResponse(fiber.StatusCreated, "Usuario registrado exitosamente", nil))
}

// Login godoc
// @Summary Inicio de sesión
// @Description Autentica a un usuario y devuelve un token JWT válido por 24h.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dtos.LoginRequest true "Credenciales de acceso"
// @Success 200 {object} dtos.ApiResponse{data=dtos.AuthResponse} "Login exitoso con token JWT"
// @Failure 401 {object} dtos.ApiResponse "Credenciales inválidas"
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c fiber.Ctx) error {
	var req dtos.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dtos.NewErrorResponse(fiber.StatusBadRequest, "Datos inválidos"))
	}

	token, user, err := ctrl.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.NewErrorResponse(fiber.StatusUnauthorized, err.Error()))
	}

	return c.JSON(dtos.NewSuccessResponse(fiber.StatusOK, "Login exitoso", dtos.AuthResponse{
		Token: token,
		User:  user,
	}))
}
