package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

type ReportRequest struct {
	UserType    string
	VolunteerID uint
	StartDate   int64
	EndDate     int64
}

type HoursWorkedByNPO struct {
	NPOID       uint
	NPOName     string
	HoursWorked int64
}

func GetVolunteerHours(ctx iris.Context) {
	db, err := databaseConfig.DbStart()
	if err != nil {
		ctx.Values().Set("message", "Unable to alter start or end times once an event has already started.")
		ctx.StatusCode(500)
	}
	defer db.Close()

	//Read request body
	var requestBody ReportRequest
	ctx.ReadJSON(&requestBody)

	//Get all Volunteer Shifts whose start dates are within the requested start and and dates
	var volunteerShifts []types.Shift
	db.Table("shifts").Where("shifts.volunteer_id = ? AND shifts.actual_start_time >= ? AND shifts.actual_start_time <= ?", requestBody.VolunteerID, requestBody.StartDate, requestBody.EndDate).Find(&volunteerShifts)

	//Get Hours worked by volunteer sorted by NPO
	var hoursWorkedByNPO map[uint]HoursWorkedByNPO
	for _, shift := range volunteerShifts {
		var eventInfo types.Event
		db.First(&eventInfo, shift.EventID)
		_, ok := hoursWorkedByNPO[eventInfo.NPOID]
		if ok {

			var npoInfo types.NPO
			db.Select("npo_name").First(&npoInfo, eventInfo.NPOID)
			shiftMinutes := (shift.ActualEndTime - shift.ActualStartTime) / 60000
			shiftHours := shiftMinutes / 60
			updatedShiftHours := hoursWorkedByNPO[eventInfo.NPOID].HoursWorked + shiftHours
			npoHoursWorked := HoursWorkedByNPO{
				NPOID:       eventInfo.NPOID,
				NPOName:     npoInfo.NPOName,
				HoursWorked: updatedShiftHours,
			}
			hoursWorkedByNPO[eventInfo.NPOID] = npoHoursWorked

		} else {
			var npoInfo types.NPO
			db.Select("npo_name").First(&npoInfo, eventInfo.NPOID)
			shiftMinutes := (shift.ActualEndTime - shift.ActualStartTime) / 60000
			shiftHours := shiftMinutes / 60
			npoHoursWorked := HoursWorkedByNPO{
				NPOID:       eventInfo.NPOID,
				NPOName:     npoInfo.NPOName,
				HoursWorked: shiftHours,
			}
			hoursWorkedByNPO[eventInfo.NPOID] = npoHoursWorked
		}
	}
	fmt.Print(hoursWorkedByNPO)
	ctx.JSON(hoursWorkedByNPO)

}
