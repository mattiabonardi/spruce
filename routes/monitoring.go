package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create customer
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
