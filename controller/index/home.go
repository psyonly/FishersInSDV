package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EntryPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome 2 FishersInSDV",
	})
}
