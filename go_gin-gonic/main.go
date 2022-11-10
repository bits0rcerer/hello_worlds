package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// create gin's default engine
	engine := gin.Default()

	// apply our routes
	applyRoutes(engine)

	// listen on port 8080
	log.Panicln(engine.Run(":8080"))
}
