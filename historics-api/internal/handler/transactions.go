package handler

import (
	"github.com/gin-gonic/gin"
	"historics-api/internal/dto"
	"net/http"
	"strconv"
)

func (h *Handler) GetTransactions(c *gin.Context) {
	collection := c.Param("collection")
	l := c.Query("limit")
	p := c.Query("page")

	var req = dto.GetTransaction{
		Collection: collection,
	}

	if l == "" {
		req.Limit = 25
	} else {
		limit, _ := strconv.ParseInt(l, 10, 64)
		req.Limit = limit
	}

	if p == "" {
		req.Page = 0
	} else {
		page, _ := strconv.ParseInt(p, 10, 64)
		req.Page = page
	}

	tx, err := h.TransactionService.GetTransactions(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
	}

	if tx == nil {
		c.JSON(http.StatusOK, gin.H{
			"transactions": nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":         req.Page,
		"limit":        req.Limit,
		"transactions": tx,
	})
}
