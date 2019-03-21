package controllers

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

func GetAllEvents(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var events []types.Event

	db.Find(&events)

	ctx.JSON(events)
}

func ShowEvent(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var event types.Event

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&event, urlParam)

	db.Model(&event).Related(&event.Shifts)

	ctx.JSON(event)
}

func CreateEvent(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody types.Event

	ctx.ReadJSON(&requestBody)

	event := types.Event{
		NPOID:           requestBody.NPOID,
		Name:            requestBody.Name,
		StartTime:       requestBody.StartTime,
		EndTime:         requestBody.EndTime,
		Tags:            requestBody.Tags,
		Description:     requestBody.Description,
		Location:        requestBody.Location,
		NumOfVolunteers: requestBody.NumOfVolunteers,
	}

	db.NewRecord(event)
	db.Create(&event)

	for i := 0; i < event.NumOfVolunteers; i++ {
		shift := types.Shift{
			EventID:         event.ID,
			ActualStartTime: event.StartTime,
			ActualEndTime:   event.EndTime,
		}
		db.NewRecord(shift)
		db.Create(&shift)
	}

	var newEvent types.Event

	db.First(&newEvent, event.ID)

	db.Model(&event).Related(&newEvent.Shifts)

	ctx.JSON(newEvent)
}

func UpdateEvent(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var event types.Event

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&event, urlParam)

	var requestBody types.Event

	ctx.ReadJSON(&requestBody)

	now := time.Now().Unix()

	//if the event hasn't started:
	if now < event.StartTime {

		// if the NPO changes the start time and the event hasn't started, also change the start time for all shifts
		if requestBody.StartTime != event.StartTime {
			fmt.Print("The request start time is not equal to the event start time")
			db.Table("shifts").Where("event_id = ?", event.ID).Updates(map[string]interface{}{"actual_start_time": requestBody.StartTime})
		}

		//if the NPO changes the end time, and the event hasn't started, also change the end time for all shifts
		if requestBody.EndTime != event.EndTime {
			db.Table("shifts").Where("event_id = ?", event.ID).Updates(map[string]interface{}{"actual_end_time": requestBody.EndTime})
		}

		//change updated fields on Event itself including start/end times
		db.Model(&event).Updates(types.Event{
			Name:        requestBody.Name,
			StartTime:   requestBody.StartTime,
			EndTime:     requestBody.EndTime,
			Tags:        requestBody.Tags,
			Description: requestBody.Description,
			Location:    requestBody.Location,
		})

		var newEvent types.Event

		db.First(&newEvent, event.ID)

		db.Model(&event).Related(&newEvent.Shifts)

		ctx.JSON(newEvent)

	} else if event.StartTime == requestBody.StartTime && event.EndTime == requestBody.EndTime {
		//change updated fields on Event itself including start/end times
		db.Model(&event).Updates(types.Event{
			Name:        requestBody.Name,
			StartTime:   requestBody.StartTime,
			EndTime:     requestBody.EndTime,
			Tags:        requestBody.Tags,
			Description: requestBody.Description,
			Location:    requestBody.Location,
		})

		var newEvent types.Event

		db.First(&newEvent, event.ID)

		db.Model(&event).Related(&newEvent.Shifts)

		ctx.JSON(newEvent)
	} else {
		ctx.Values().Set("message", "Unable to alter start or end times once an event has already started.")
		ctx.StatusCode(500)
	}

}
