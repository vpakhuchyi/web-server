package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/vpakhuchyi/web-server/routers"
)

func main() {

	r := gin.Default()
	t, _ := template.ParseFiles("templates/default.tmpl")
	r.SetHTMLTemplate(t)
	r.GET("/searchText", routers.GETJSONHandler)
	r.POST("/searchText", routers.POSTJSONHandler)
	r.Run(":8080")

}
