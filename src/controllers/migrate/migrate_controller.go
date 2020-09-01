package migrate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rampo0/multi-lang-microservice/users/src/datasources/mysql/users_db/migrations"
)

func Migrate(c *gin.Context) {
	migrations.User()
	c.String(http.StatusOK, "migrate success")
}
