package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superdima3000/transport-auth/db"
)

// @Summary      Создание карты
// @Description  Добавляет новую карту в систему
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body db.Card true "Данные карты"
// @Success      200  {object}  map[string]int "ID созданной карты"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /cards [post]
func (h *Handler) createCard(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	var input db.Card
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	id, err := h.services.Card.Create(input)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      Получение всех карт
// @Description  Возвращает список всех карт
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   db.Card "Список карт"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /cards [get]
func (h *Handler) getAllCards(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}
	cards, err := h.services.Card.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

// @Summary      Получение карты по ID
// @Description  Возвращает данные карты по идентификатору
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID карты"
// @Success      200  {object}  db.Card "Данные карты"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /cards/{id} [get]
func (h *Handler) getCardById(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	card, err := h.services.Card.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, card)
}

// @Summary      Обновление карты
// @Description  Обновляет данные существующей карты
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID карты"
// @Param        input body db.UpdateCard true "Данные для обновления"
// @Success      204  "Успешно обновлено"
// @Failure      400  {object}  map[string]string "Неверные параметры или ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /cards/{id} [put]
func (h *Handler) updateCard(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	var input db.UpdateCard
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	err = h.services.Card.Update(id, input)
	if err != nil {
		handleServiceError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// @Summary      Удаление карты
// @Description  Удаляет карту по ID
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID карты"
// @Success      204  "Успешно удалено"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /cards/{id} [delete]
func (h *Handler) deleteCard(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	err = h.services.Card.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
