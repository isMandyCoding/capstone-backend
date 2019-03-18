package main

import (
	"fmt"

	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/theycallmethetailor/capstone-backend/controllers"

	types "github.com/theycallmethetailor/capstone-backend/models"

	_ "github.com/lib/pq"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	db, _ := databaseConfig.DbStart()

	db.AutoMigrate(&types.NPO{}, &types.Volunteer{}, &types.Opportunity{}, &types.Fulfiller{})

	fmt.Println("Works")

	app.Get("/api/npos", controllers.GetAllNPOs)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
