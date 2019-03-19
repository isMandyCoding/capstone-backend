package controllers

import (
	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

func GetAllOpportunities(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var opportunities []types.Opportunity

	db.Find(&opportunities)


	ctx.JSON(opportunities)
}

func ShowOpportunity(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var opportunity types.Opportunity

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&opportunity, urlParam)

	db.Model(&opportunity).Related(&opportunity.Fulfillers)

	ctx.JSON(opportunity)
}

func CreateOpportunity(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody types.Opportunity

	ctx.ReadJSON(&requestBody)

	opportunity := types.Opportunity{
		NPOID:           requestBody.NPOID,
		StartTime:       requestBody.StartTime,
		EndTime:         requestBody.EndTime,
		Label:           requestBody.Label,
		Tags:            requestBody.Tags,
		Description:     requestBody.Description,
		Location:        requestBody.Location,
		NumOfVolunteers: requestBody.NumOfVolunteers,
	}

	db.NewRecord(opportunity)
	db.Create(&opportunity)

	if db.NewRecord(opportunity) == false {
		ctx.JSON(opportunity)
	}
}

func DeleteOpportunity(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var opportunity types.Opportunity
	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&opportunity, urlParam)

	db.Unscoped().Delete(&opportunity)

	ctx.JSON(opportunity)

}


