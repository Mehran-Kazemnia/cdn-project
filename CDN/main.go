package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	router.GET("/cdn", func(c *gin.Context) {
		dt_str := time.Now().UTC().String()

		c.HTML(http.StatusOK, "cdn.html", gin.H{
			"time": dt_str,
		})
	})
	router.GET("/host", func(c *gin.Context) {
		dt_str := time.Now().UTC().String()
		resp, err := http.Get("http://localhost:8090")
		defer resp.Body.Close()
		responseData, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}
		c.HTML(http.StatusOK, "host.html", gin.H{
			"sendTime":    dt_str,
			"receiveTime": string(responseData),
		})
	})
	router.Run("localhost:8080")
}
