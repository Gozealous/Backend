package nexus

import (
	"net/http"

	"gozealous/service/health"

	"github.com/gin-gonic/gin"
)

func HandleHealthStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		result := health.Status()

		response := map[string]string{
			"status": "success",
			"data":   result,
		}

		c.JSON(http.StatusOK, response)
	}
}
