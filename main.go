package main

import (
	"fmt"

	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	controllers "github.com/theycallmethetailor/capstone-backend/controllers"

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

	// NPO Routes:
	app.Get("/api/npos", controllers.GetAllNPOs)
	app.Get("/api/npos/{id:int}", controllers.ShowNPO)
	app.Post("api/npos", controllers.CreateNPO)
	app.Put("/api/npos/{id:int}", controllers.UpdateNPO)
	app.Delete("api/npos/{id:int}", controllers.DeleteNPO)

	//Volunteers Routes:
	app.Get("/api/volunteers", controllers.GetAllVolunteers)
	app.Get("/api/volunteers/{id:int}", controllers.ShowVolunteer)
	app.Post("/api/volunteers", controllers.CreateVolunteer)
	app.Put("/api/volunteers/{id:int}", controllers.UpdateVolunteer)
	app.Delete("/api/volunteers/{id:int}", controllers.DeleteVolunteer)

	//Opportunities Routes:
	app.Get("/api/opportunities", controllers.GetAllOpportunities)
	app.Get("/api/opportunities/{id:int}", controllers.ShowOpportunity)
	app.Post("api/opportunities", controllers.CreateOpportunity)
	app.Delete("/api/opportunities/{id:int}", controllers.DeleteOpportunity)

	//Fulfillers Routes:
	app.Get("/api/fulfillers/opportunty/{oppid:int}", controllers.GetOppFulfillers) //get by opp id
	app.Get("/api/fulfullers/volunteer/{volid:int}", controllers.GetVolFulfillers) //get by vol id
	app.Get("/api/fulfillers", controllers.GetAllFulfillers)
	app.Get("/api/fulfillers/{id:int}", controllers.ShowFulfiller)
	app.Post("/api/fulfillers", controllers.CreateFulfiller)
	app.Put("api/fulfillers/{id:int}", controllers.UpdateFulfiller)
	app.Delete("/api/fulfillers/{id:int}", controllers.DeleteFulfiller)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
