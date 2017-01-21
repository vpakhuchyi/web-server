package main

import (
	"fmt"

	"github.com/vpakhuchyi/web-server/controllers"
)

func main() {
	// r := gin.Default()
	// t, _ := template.ParseFiles("templates/default.tmpl")
	// r.SetHTMLTemplate(t)
	// r.GET("/searchText", controllers.GETJSONHandler)
	// r.POST("/searchText", controllers.POSTJSONHandler)
	// r.Run(":8080")

	sites := make([]string, 2)
	sites[0] = "https://google.com.ua"
	sites[1] = "https://yahoo.com"
	res, err := controllers.SearchForArgOnSites("Google", sites)
	fmt.Println(res, err)
}
