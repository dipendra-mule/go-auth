package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"authentication/config"
	"authentication/helpers"
	"authentication/routes"
)

func main() {
	 // log in to mongoDB

	 log.Println("Starting applocation...")
	  key := config.GenerateRandomKey()
	  helpers.SetJWTKey(key)
	  fmt.Println("Generated JWT Key: ", key)
	  // Init gin router

	  r := gin.Default()
	  routes.SetupRoutes(r)

	  // Start the server

	  r.Run(":8080")
	  log.Println("Server is runinng on https://localhost:8080")
}