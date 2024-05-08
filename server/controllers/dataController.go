package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goameer030/732121104037/server/data"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetProductsOnCategory(c *gin.Context) {
	category := c.Param("category")
	found := false
	for _, v := range data.CATEGORIES {
		if v == category {
			product, err := GetProducts("AMZ", category)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error getting products",
				})
			} else {
				c.JSON(http.StatusOK, gin.Bind(product))
			}
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Category not found",
		})
	}
}
