package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/service"
)

// @Summary      Создание транзакции
// @Description  Создает новую транзакцию
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body db.Transaction true "Данные транзакции"
// @Success      200  {object}  map[string]int "ID транзакции"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /transactions [post]
func (h *Handler) createTransaction(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	var input db.Transaction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	id, err := h.services.Transaction.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      Авторизация транзакции
// @Description  Пытается провести авторизацию транзакции по ID
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID транзакции"
// @Success      200  {object}  map[string]string "Статус операции (approved/declined)"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /transactions/{id}/authorize [post]
func (h *Handler) authorizeTransaction(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	status, err := h.services.Transaction.Authorize(id)

	if err != nil {
		if errors.Is(err, service.ErrNotEnoughFund) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "declined",
				"message": err.Error(),
			})
		} else {
			handleServiceError(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

// @Summary      Получение всех транзакций
// @Description  Возвращает список всех транзакций
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   db.Transaction "Список транзакций"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /transactions [get]
func (h *Handler) getAllTransactions(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}
	transactions, err := h.services.Transaction.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// @Summary      Получение транзакции по ID
// @Description  Возвращает данные транзакции по идентификатору
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID транзакции"
// @Success      200  {object}  db.Transaction "Данные транзакции"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /transactions/{id} [get]
func (h *Handler) getTransactionById(c *gin.Context) {
	_, err := getUserIDAndRole(c, false)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	transaction, err := h.services.Transaction.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// @Summary      Обновление транзакции
// @Description  Обновляет данные транзакции
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID транзакции"
// @Param        input body db.UpdateTransaction true "Данные для обновления"
// @Success      204  "Успешно обновлено"
// @Failure      400  {object}  map[string]string "Неверные параметры"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /transactions/{id} [put]
func (h *Handler) updateTransaction(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	var input db.UpdateTransaction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid parameters")
		return
	}

	err = h.services.Transaction.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// @Summary      Удаление транзакции
// @Description  Удаляет транзакцию по ID
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID транзакции"
// @Success      204  "Успешно удалено"
// @Failure      400  {object}  map[string]string "Неверный ID"
// @Failure      401  {object}  map[string]string "Не авторизован"
// @Failure      403  {object}  map[string]string "Нет доступа"
// @Failure      500  {object}  map[string]string "Ошибка сервера"
// @Router       /transactions/{id} [delete]
func (h *Handler) deleteTransaction(c *gin.Context) {
	_, err := getUserIDAndRole(c, true)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id")
		return
	}

	err = h.services.Transaction.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
