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
	urlParam, _ := ctx.Params().GetInt("shiftid")
	db.First(&shift, urlParam)

	type ShiftVolunteer struct {
		VolunteerID uint
	}
	var shiftVol ShiftVolunteer
	ctx.ReadJSON(&shiftVol)

	//check to make sure volunteer hasn't already signed up for a shift for the same event

	var checkShift []types.Shift
	signup := db.Find(&checkShift, "event_id = ? AND volunteer_id = ?", shift.EventID, shiftVol.VolunteerID)

	if signup.RowsAffected == 0 {

		//check to make sure whey don't have any existing shifts during the same time
		var priorObligations []types.Shift
		obligations := db.Find(&priorObligations, "volunteer_id = ? AND actual_start_time <= ? AND actual_end_time >= ?", shiftVol.VolunteerID, shift.ActualStartTime, shift.ActualStartTime)

		if obligations.RowsAffected == 0 {

			var updatedShift types.Shift
			db.First(&updatedShift, urlParam)
			db.Model(&updatedShift).Updates(types.Shift{
				VolunteerID: shiftVol.VolunteerID,
			})
			ctx.JSON(updatedShift)

		} else {

			ctx.Values().Set("message", "You've already signed up for another event during the same time.")
			ctx.StatusCode(500)
		}

	} else {

		ctx.Values().Set("message", "You've already signed up for this event. Unable to sign up for another shift for same event.")
		ctx.StatusCode(500)
	}
}

func VolunteerCancel(ctx iris.Context) {
	db, err := databaseConfig.DbStart()

	defer db.Close()

	if err != nil {
		ctx.Values().Set("message", "Unable to update shift as requested. Please try again.")
		ctx.StatusCode(500)
	}

	defer db.Close()
	var shift types.Shift
	urlParam, _ := ctx.Params().GetInt("shiftid")
	db.First(&shift, urlParam)

	//double check to make sure the correct volunteer is being removed
	if int(shift.VolunteerID) == urlParam {

		shift.VolunteerID = 0
		db.Save(&shift)

		ctx.JSON(shift)
	}

}

func GetVolunteerShifts(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()
	defer db.Close()

	var shifts []types.Shift

	urlParam, _ := ctx.Params().GetInt("id")

	db.Where("volunteer_id = ?", urlParam).Find(&shifts)

	ctx.JSON(shifts)

}
