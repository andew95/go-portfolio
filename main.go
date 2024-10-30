package main

import (
	"fmt"
	"portfolio/app"
	"portfolio/config"
	"portfolio/db"
	"portfolio/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.NewConfig()
	fmt.Println(conf)

	// config port & adapter
	db := db.Init(conf.PostgreSQL.BaseUrl, conf.PostgreSQL.Username, conf.PostgreSQL.Password, conf.PostgreSQL.DbName)
	repo := app.SetupRepository(db)
	serv := app.SetupAppSevice(conf, repo)
	ctl := app.SetupHandler(conf, serv)

	r := gin.Default()

	router.SetupRouter(r, ctl)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
