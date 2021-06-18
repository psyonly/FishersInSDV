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

func GetFishes(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, []gin.H{
		{
			"name":   "Anchovy",
			"price":  35,
			"region": "beach",
		},
		{
			"name":   "Dorado",
			"price":  120,
			"region": "coal woods",
		},
		{
			"name":   "Tuna",
			"price":  70,
			"region": "pier",
		},
		{
			"name":   "Octopus",
			"price":  180,
			"region": "deep sea",
		},
	})
}
