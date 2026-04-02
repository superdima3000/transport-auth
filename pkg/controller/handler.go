package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/superdima3000/transport-auth/docs"
	"github.com/superdima3000/transport-auth/pkg/service"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api", h.userIdentity)
	{
		v1 := api.Group("/v1")
		{
			terminals := v1.Group("/terminals")
			{
				terminals.GET("/", h.getAllTerminals)
				terminals.POST("/", h.createTerminal)
				terminals.GET("/:id", h.getTerminalById)
				terminals.PUT("/:id", h.updateTerminal)
				terminals.DELETE("/:id", h.deleteTerminal)
			}

			cards := v1.Group("/cards")
			{
				cards.GET("/", h.getAllCards)
				cards.POST("/", h.createCard)
				cards.GET("/:id", h.getCardById)
				cards.PUT("/:id", h.updateCard)
				cards.DELETE("/:id", h.deleteCard)
			}

			keys := v1.Group("/keys")
			{
				keys.GET("/", h.getAllKeys)
				keys.POST("/", h.createKey)
				keys.GET("/:id", h.getKeyById)
				keys.PUT("/:id", h.updateKey)
				keys.DELETE("/:id", h.deleteKey)
			}

			transactions := v1.Group("/transactions")
			{
				transactions.GET("/", h.getAllTransactions)
				transactions.POST("/", h.createTransaction)
				transactions.GET("/:id", h.getTransactionById)
				transactions.PUT("/:id", h.updateTransaction)
				transactions.DELETE("/:id", h.deleteTransaction)
				transactions.GET("/:id/authorize", h.authorizeTransaction)
			}
		}
	}

	return router
}

func handleServiceError(c *gin.Context, err error) {
	if appErr, ok := errors.AsType[*service.AppError](err); ok {
		newErrorResponse(c, appErr.StatusCode, appErr.Message)
		return
	}
	newErrorResponse(c, http.StatusInternalServerError, err.Error())
}
