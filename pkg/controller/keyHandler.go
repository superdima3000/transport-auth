package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superdima3000/transport-auth/db"
)

// @Summary      Создание ключа
// @Description  Добавляет новый ключ доступа
// @Tags         keys
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body db.Key true "Данные ключа"
// @Success      200  {object}  map[string]int "ID ключа"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /keys [post]
func (h *Handler) createKey(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	var input db.Key
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	id, err := h.services.Key.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      Получение всех ключей
// @Description  Возвращает список всех ключей
// @Tags         keys
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   db.Key "Список ключей"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /keys [get]
func (h *Handler) getAllKeys(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	keys, err := h.services.Key.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": keys})
}

// @Summary      Получение ключа по ID
// @Description  Возвращает данные ключа по идентификатору
// @Tags         keys
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID ключа"
// @Success      200  {object}  db.Key "Данные ключа"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /keys/{id} [get]
func (h *Handler) getKeyById(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	key, err := h.services.Key.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, key)
}

// @Summary      Обновление ключа
// @Description  Обновляет данные ключа
// @Tags         keys
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID ключа"
// @Param        input body db.UpdateKey true "Данные для обновления"
// @Success      204  "Успешно обновлено"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /keys/{id} [put]
func (h *Handler) updateKey(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	var input db.UpdateKey
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	err = h.services.Key.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// @Summary      Удаление ключа
// @Description  Удаляет ключ по ID
// @Tags         keys
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID ключа"
// @Success      204  "Успешно удалено"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /keys/{id} [delete]
func (h *Handler) deleteKey(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	err = h.services.Key.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
