package main

import (
	"os"

	config "api-sql/config"
	"api-sql/routes"

	"github.com/gin-gonic/gin"
)

var DBUSER string = os.Getenv("DBUSER")
var DBPASS string = ""
var DBNAME string = "recordings"

func main() {
	//Connect to the database
	config.ConnectToDatabase(DBUSER, DBPASS, DBNAME)

	router := gin.Default()

	routes.RouteAlbums(router)

	router.Run("localhost:8080")

}
