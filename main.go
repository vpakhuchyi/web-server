package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/vpakhuchyi/geekstest/controllers"
)

func main() {
	r := gin.Default()
	t, _ := template.ParseFiles("templates/default.tmpl")
	r.SetHTMLTemplate(t)
	r.GET("/searchText", controllers.GETJSONHandler)
	r.POST("/searchText", controllers.POSTJSONHandler)
	r.Run(":8080")
}
