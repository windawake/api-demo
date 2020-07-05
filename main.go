package main

import (
	"api-demo/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// gin.DisableConsoleColor()

	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	routes.InitRouter().Run() // listen and serve on 0.0.0.0:8080

}
