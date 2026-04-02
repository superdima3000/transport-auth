package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superdima3000/transport-auth/db"
)

// @Summary      Создание пользователя
// @Description  Создает нового пользователя (требуется роль администратора)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body db.User true "Данные пользователя (login, password, role)"
// @Success      200  {object}  map[string]int "ID созданного пользователя"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Недостаточно прав (требуется администратор)"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /users [post]
func (h *Handler) createUser(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	var input db.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	id, err := h.services.User.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      Получение всех пользователей
// @Description  Возвращает список всех пользователей (доступно всем авторизованным)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{} "Список пользователей в поле data"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}
	users, err := h.services.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// @Summary      Получение пользователя по ID
// @Description  Возвращает данные пользователя по идентификатору (доступно всем авторизованным)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID пользователя"
// @Success      200  {object}  db.User "Данные пользователя"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	user, err := h.services.User.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Обновление пользователя
// @Description  Обновляет данные пользователя (требуется роль администратора)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID пользователя"
// @Param        input body db.UpdateUser true "Данные для обновления"
// @Success      204  "Успешно обновлено"
// @Failure      400  {object}  map[string]string "Неверные параметры или ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Недостаточно прав (требуется администратор)"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	var input db.UpdateUser
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	err = h.services.User.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// @Summary      Удаление пользователя
// @Description  Удаляет пользователя по ID (требуется роль администратора)
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID пользователя"
// @Success      204  "Успешно удалено"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Недостаточно прав (требуется администратор)"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	err = h.services.User.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
