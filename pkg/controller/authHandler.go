package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/service"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary      Регистрация пользователя
// @Description  Создает нового пользователя в системе
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body db.User true "Данные пользователя (login, password, is_admin)"
// @Success      200  {object}  map[string]int "Возвращает ID созданного пользователя"
// @Failure      400  {object}  map[string]string "Ошибка валидации"
// @Failure      500  {object}  map[string]string "Внутренняя ошибка сервера"
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input db.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		newErrorResponse(c, http.StatusConflict, "user already exists") // 409
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      Авторизация пользователя
// @Description  Ввод логина и пароля для получения токена доступа
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body signInInput true "Логин и пароль"
// @Success      200  {object}  map[string]string "Возвращает JWT токен"
// @Failure      400  {object}  map[string]string "Ошибка валидации"
// @Failure      500  {object}  map[string]string "Ошибка авторизации"
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if errors.Is(err, service.ErrUserNotFound) {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}
	if errors.Is(err, service.ErrInvalidPassword) {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
