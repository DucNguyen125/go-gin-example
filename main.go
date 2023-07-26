package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	mysql_util "example/utils"

	"example/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("RUN_MODE"))
	mysql_util.Connect()
	routersInit := routers.InitRouter()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	endPoint := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
