package main

import (
	"github.com/iris-contrib/clean-architecture/config"
	"github.com/iris-contrib/clean-architecture/infrastructure/datastore"
	"github.com/iris-contrib/clean-architecture/infrastructure/router"
	"github.com/iris-contrib/clean-architecture/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	app := iris.New()
	router.NewRouter(app, r.NewAppController())

	addr := "localhost:" + config.C.Server.Address

	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithSitemap("http://"+addr))
}
