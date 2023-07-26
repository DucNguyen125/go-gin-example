package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"example/utils/goth"
	"example/utils/logger"
	"example/utils/mysql"

	"example/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("RUN_MODE"))

	// Init logger
	logger.Init()

	// Connect to database
	if err = mysql.Connect(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Migrate database
	if err = mysql.AutoMigrate(); err != nil {
		log.Fatal("Error migrating to database")
	}

	// Config goth
	goth.Init()
	// Register Router
	routersInit := routers.InitRouter()

	var port int
	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error get PORT")
	}
	endPoint := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:              endPoint,
		ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd // common
		Handler:           routersInit,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("Fail to start error server")
	}
}
