package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userID              = "userID"
	isAdmin             = "isAdmin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "no authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	claims, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	c.Set("userID", claims.UserID)
	c.Set("isAdmin", claims.IsAdmin)

}

func getUserIDAndRole(c *gin.Context, admin bool) (int64, error) {
	id, ok := c.Get(userID)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id not found")
		return 0, fmt.Errorf("user not found")
	}

	isAdminVal, _ := c.Get(isAdmin)
	isAdminBool, ok := isAdminVal.(int64)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user isAdmin not int")
		return 0, fmt.Errorf("invalid isAdmin type")
	}

	if admin && isAdminBool == 0 {
		newErrorResponse(c, http.StatusForbidden, "not enough rights")
		return 0, fmt.Errorf("not enough rights")
	}

	return id.(int64), nil
}
