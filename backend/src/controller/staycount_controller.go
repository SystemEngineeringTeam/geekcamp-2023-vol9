package controller

import(
	"github.com/gin-gonic/gin"
)

func StayCount(c *gin.Context){

	location_key := c.Param("location_key")

	c.JSON(200, gin.H{
		"type": "succeeded",
		"location_key": location_key,
	})
}