package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	helpers "gin/helpers"
)

func IndexGetHandler() gin.HandlerFunc {
	items := helpers.GetAllItems()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	}
}

func CreateItemHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		item := c.PostForm("item")
		result := helpers.CreateOneItem(item)
		if !result {
			log.Fatal("faild to create item...")
		}
		// c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
		// c.Redirect(http.StatusFound, "/api/")
		items := helpers.GetAllItems()
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	}
}