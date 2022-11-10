package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	applyRoutes(engine)
	log.Panicln(engine.Run(":8080"))
}
