package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superdima3000/transport-auth/db"
)

// @Summary      Создание терминала
// @Description  Добавляет новый терминал
// @Tags         terminals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body db.Terminal true "Данные терминала"
// @Success      200  {object}  map[string]int "ID терминала"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /terminals [post]
func (h *Handler) createTerminal(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}

	var input db.Terminal

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	id, err := h.services.Terminal.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      Получение всех терминалов
// @Description  Возвращает список всех терминалов
// @Tags         terminals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   db.Terminal "Список терминалов"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /terminals [get]
func (h *Handler) getAllTerminals(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}

	terminals, err := h.services.Terminal.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": terminals})
}

// @Summary      Получение терминала по ID
// @Description  Возвращает данные терминала по идентификатору
// @Tags         terminals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID терминала"
// @Success      200  {object}  db.Terminal "Данные терминала"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /terminals/{id} [get]
func (h *Handler) getTerminalById(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	terminal, err := h.services.Terminal.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, terminal)
}

// @Summary      Обновление терминала
// @Description  Обновляет данные терминала
// @Tags         terminals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID терминала"
// @Param        input body db.UpdateTerminal true "Данные для обновления"
// @Success      204  "Успешно обновлено"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /terminals/{id} [put]
func (h *Handler) updateTerminal(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	var input db.UpdateTerminal
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
	}

	err = h.services.Terminal.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// @Summary      Удаление терминала
// @Description  Удаляет терминал по ID
// @Tags         terminals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID терминала"
// @Success      204  "Успешно удалено"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /terminals/{id} [delete]
func (h *Handler) deleteTerminal(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	err = h.services.Terminal.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}
