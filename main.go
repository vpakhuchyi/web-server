package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vpakhuchyi/web-server/routers"
)

func main() {
	r := gin.Default()
	r.POST("/searchText", routers.POSTJSONHandler)
	r.Run(":8080")
}
