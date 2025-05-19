package main

import (
	"MTBS/db"

	"MTBS/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDB()

	db.Collections()
	routes.SettingUpRoutes(r)

	r.Run()
}
