package app

import (
	"multi-lang-microservice/users/src/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("About to start application ...")
	router.Run(":8080")
}
