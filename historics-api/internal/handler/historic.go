package handler

import (
	"github.com/gin-gonic/gin"
	"historics-api/internal/dto"
	"net/http"
	"strconv"
)

func (h Handler) GetHistoric(c *gin.Context) {
	collection := c.Param("collection")

	var req = dto.GetHistorics{
		Collection: collection,
	}

	l := c.Query("limit")
	if l == "" {
		req.Limit = 7
	} else {
		limit, _ := strconv.ParseInt(l, 10, 64)
		req.Limit = limit
	}

	data, err := h.HistoricService.GetHistoricData(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err": "something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"limit":    req.Limit,
		"historic": data,
	})
	return

}
