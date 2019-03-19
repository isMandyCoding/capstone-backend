package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

func GetAllVolunteers(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var volunteers types.Volunteer

	db.Find(&volunteers)

	ctx.JSON(volunteers)
}

func ShowVolunteer(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var volunteer []types.Volunteer

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&volunteer, urlParam)

	ctx.JSON(volunteer)
}

func CreateVolunteer(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody types.Volunteer

	ctx.ReadJSON(&requestBody)

	fmt.Print(requestBody.Email)

	volunteer := types.Volunteer{
		Username:  requestBody.Username,
		Bio:       requestBody.Bio,
		Email:     requestBody.Email,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Password:  requestBody.Password,
	}

	db.NewRecord(volunteer)
	db.Create(&volunteer)

	if db.NewRecord(volunteer) == false {
		var newVolunteer types.Volunteer

		db.Where("email = ?", volunteer.Email).Find(&newVolunteer)

		ctx.JSON(newVolunteer)
	} else {

	}

}

func UpdateVolunteer(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var volunteer types.Volunteer

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&volunteer, urlParam)

	var requestBody types.Volunteer

	ctx.ReadJSON(&requestBody)

	db.Model(&volunteer).Updates(types.Volunteer{
		Username:  requestBody.Username,
		Bio:       requestBody.Bio,
		Email:     requestBody.Email,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Password:  requestBody.Password,
	})


	ctx.JSON(volunteer)
}

func DeleteVolunteer(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var volunteer types.Volunteer

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&volunteer, urlParam)

	var deletedVolunteer types.Volunteer

	db.Unscoped().Delete(&volunteer)

	db.First(&deletedVolunteer, urlParam)

	ctx.JSON(deletedVolunteer)
}
