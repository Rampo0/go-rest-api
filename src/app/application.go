package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rampo0/multi-lang-microservice/users/src/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("About to start application ...")
	router.Run(":8080")
}
