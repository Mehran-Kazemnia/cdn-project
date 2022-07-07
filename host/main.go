package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 2)
		dt_str := time.Now().UTC().String()
		c.String(http.StatusOK, dt_str)

	})

	router.Run("localhost:8090")
}
