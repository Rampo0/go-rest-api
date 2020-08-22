package migrate

import (
	"multi-lang-microservice/users/src/datasources/mysql/users_db/migrations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {
	migrations.User()
	c.String(http.StatusOK, "migrate success")
}
