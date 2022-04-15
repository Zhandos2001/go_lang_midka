package controller

import (
	m "midterm-project/map"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appController struct{}

type AppController interface {
	HandleBaseRoute(c *gin.Context)
	SetValue(c *gin.Context) 
	GetValue(c *gin.Context) 
}

type PutRequest struct {
	Key 	 string `json:"key"`
	Value  string `json:"value"`
}

func GetAppController() AppController {
	return &appController{}
}

func (*appController) HandleBaseRoute(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}

func (*appController) GetValue(c *gin.Context) {
	key := c.Param("key")
	value, ok := m.Map.Get(key)

	if ok {
		c.JSON(http.StatusOK, gin.H{
			"value": value,
		})	
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Key not found in map",
	})
}

func (*appController) SetValue(c *gin.Context) {
	var pair PutRequest
	c.BindJSON(&pair)

	if pair.Key == "" || pair.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Key or value is empty!",
		})
		return	
	}

	m.Map.Set(pair.Key, pair.Value)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully set new pair with key: " + pair.Key + " and value: " + pair.Value,
	})
}