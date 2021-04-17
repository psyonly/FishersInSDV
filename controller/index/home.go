package home

import (
	"net/http"
)

import (
	"github.com/psyonly/FishersInSDV/settings"
)

import (
	"github.com/gin-gonic/gin"
)

func EntryPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "welcome 2 FishersInSDV",
	})
}

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": settings.Conf.Version,
	})
}
