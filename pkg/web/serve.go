package web

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Serve serves the web server.
func Serve() {
	router := gin.New()
	router.LoadHTMLGlob("/app/html/*.html")
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Static("/static", "/app/html/static")
	router.GET("/", rootHandler)
}

// rootHandler handles the root request.
func rootHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles("/app/html/index.html")
	if err != nil {
		logrus.Error(err.Error())
		c.String(500, err.Error())
		return
	}
	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		logrus.Error(err.Error())
		c.String(500, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
