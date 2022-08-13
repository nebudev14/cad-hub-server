package main

import (
	"fmt"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/NebuDev14/cad-hub-server/db"
	"github.com/NebuDev14/cad-hub-server/routes"
	// "net/http"
)

func main() {

	router := gin.Default()

	fmt.Println("init")
	db.InitConnection()
	routes.StartChanges(router.Group("/changes"))
	routes.StartUpdates(router.Group("/updates"))
	routes.StartRegister(router.Group("/register"))

	router.Run()
}
