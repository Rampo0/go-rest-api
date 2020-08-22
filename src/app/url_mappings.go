package app

import (
	"multi-lang-microservice/users/src/controllers/migrate"
	"multi-lang-microservice/users/src/controllers/ping"
	"multi-lang-microservice/users/src/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/migrate", migrate.Migrate)

	// router.GET("/users/search", users.SearchUser)
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}
