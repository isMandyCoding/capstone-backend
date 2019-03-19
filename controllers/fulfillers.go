package controllers

import (
	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)


//returns all fulfillments by opportunity id
func GetOppFulfillers(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var opportunity types.Opportunity

	urlParam, _ := ctx.Params().GetInt("oppid")

	db.First(&opportunity, urlParam)

	db.Model(&opportunity).Related(&opportunity.Fulfillers)

	ctx.JSON(opportunity.Fulfillers)

}

//returns all fulfillments by volunteer id
func GetVolFulfillers(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var fulfillers []types.Fulfiller

	urlParam, _ := ctx.Params().GetInt("volid")


	db.Where("volunteer_id = ?", urlParam).Find(&fulfillers)

	ctx.JSON(fulfillers)

}

//returns all fillfilments
func GetAllFulfillers (ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var fulfillers []types.Fulfiller

	db.Find(&fulfillers)
	
	ctx.JSON(&fulfillers)
}

func ShowFulfiller (ctx iris.Context) {
	
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var fulfiller types.Fulfiller

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&fulfiller, urlParam)

	ctx.JSON(fulfiller)
}

func CreateFulfiller (ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody types.Fulfiller

	ctx.ReadJSON(&requestBody)

	fulfiller := types.Fulfiller{
		OpportunityID: requestBody.OpportunityID,
		VolunteerID: requestBody.VolunteerID,
		Approved: requestBody.Approved,
		PlannedStart: requestBody.PlannedStart,
		PlannedEnd: requestBody.PlannedEnd,
		ActualStart: requestBody.ActualStart,
		ActualEnd: requestBody.ActualEnd,
	}

	db.NewRecord(fulfiller)
	db.Create(&fulfiller)

	if db.NewRecord(fulfiller) == false {
		ctx.JSON(fulfiller)
	}

}

func UpdateFulfiller(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var fulfiller types.Fulfiller

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&fulfiller, urlParam)

	var requestBody types.Fulfiller

	ctx.ReadJSON(&requestBody)

	db.Model(&fulfiller).Updates(types.Fulfiller{
		OpportunityID: requestBody.OpportunityID,
		VolunteerID: requestBody.VolunteerID,
		Approved: requestBody.Approved,
		PlannedStart: requestBody.PlannedStart,
		PlannedEnd: requestBody.PlannedEnd,
		ActualStart: requestBody.ActualStart,
		ActualEnd: requestBody.ActualEnd,
	})

	ctx.JSON(fulfiller)

}

func DeleteFulfiller(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var fulfiller types.Fulfiller

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&fulfiller, urlParam)

	db.Unscoped().Delete(&fulfiller)

	ctx.JSON(fulfiller)
	
}