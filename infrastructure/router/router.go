package router

import (
	"github.com/iris-contrib/clean-architecture/interface/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func NewRouter(app *iris.Application, c controller.AppController) {
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/users", func(ctx iris.Context) {
		if err := c.GetUsers(ctx); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("GetUsers: %v", err)
		}
	})

	// mvc.Configure(app, func(mvcApp *mvc.Application){
	// 	mvcApp.Register(...dependencies)
	// 	mvcApp.Handle(new(controller))
	// })
}
