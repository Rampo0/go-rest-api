package app

import (
	"github.com/rampo0/multi-lang-microservice/users/src/controllers/migrate"
	"github.com/rampo0/multi-lang-microservice/users/src/controllers/ping"
	"github.com/rampo0/multi-lang-microservice/users/src/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/migrate", migrate.Migrate)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}
