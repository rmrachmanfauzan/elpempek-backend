package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rmrachmanfauzan/elpempek/internal/dto"
	"github.com/rmrachmanfauzan/elpempek/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s}
}


func (h *UserHandler) CreateUser(c echo.Context) error {
    var req dto.CreateUserRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    user, err := h.service.CreateUser(c.Request().Context(), req)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
    }

    return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		
	}
	user, err := h.service.GetUserByID(c.Request().Context(),uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "User not found"})
		
	}
	return c.JSON(http.StatusOK, user)
}
