package main

import "github.com/gin-gonic/gin"

// main initializes the Gin router and starts the server on port 3000.

var Storage = make(map[string]string)

func main() {

	router := gin.Default()

	router.Run(":3000")
}
