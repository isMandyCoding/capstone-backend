package controllers

import (
	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

func VolunteerSignup(ctx iris.Context) {
	db, err := databaseConfig.DbStart()

	if err != nil {
		ctx.Values().Set("message", "Unable to update shift as requested. Please try again.")
		ctx.StatusCode(500)
	}

	defer db.Close()

	var shift types.Shift

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&shift, urlParam)

	type ShiftVolunteer struct {
		VolunteerID uint
	}

	var shiftVol ShiftVolunteer

	ctx.ReadJSON(&shiftVol)

	var updatedShift types.Shift

	db.First(&updatedShift, urlParam)

	db.Model(&updatedShift).Updates(types.Shift{
		VolunteerID: shiftVol.VolunteerID,
	})

	ctx.JSON(updatedShift)
}
