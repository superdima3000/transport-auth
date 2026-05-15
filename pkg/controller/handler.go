package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"}, // твой фронтенд
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
				terminals.GET("", h.getAllTerminals)
				terminals.POST("", h.createTerminal)
				terminals.GET("/:id", h.getTerminalById)
				terminals.PUT("/:id", h.updateTerminal)
				terminals.DELETE("/:id", h.deleteTerminal)
			}

			cards := v1.Group("/cards")
			{
				cards.GET("", h.getAllCards)
				cards.POST("", h.createCard)
				cards.GET("/:id", h.getCardById)
				cards.PUT("/:id", h.updateCard)
				cards.DELETE("/:id", h.deleteCard)
			}

			keys := v1.Group("/keys")
			{
				keys.GET("", h.getAllKeys)
				keys.POST("", h.createKey)
				keys.GET("/:id", h.getKeyById)
				keys.PUT("/:id", h.updateKey)
				keys.DELETE("/:id", h.deleteKey)
			}

			transactions := v1.Group("/transactions")
			{
				transactions.GET("", h.getAllTransactions)
				transactions.POST("", h.createTransaction)
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
