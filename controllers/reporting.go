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

type HoursWorkedByEvent struct {
	EventID        uint
	EventName      string
	EventStartTime int64
	EventEndTime   int64
	HoursWorked    int64
}

type HoursWorkedByTag struct {
	TagID       uint
	TagName     string
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

	//Get Hours worked sorted by NPO
	var hoursWorkedByNPO map[uint]HoursWorkedByNPO
	hoursWorkedByNPO = make(map[uint]HoursWorkedByNPO)

	//Acumulate hours worked sorted by Event
	var hoursWorkedByEvent map[uint]HoursWorkedByEvent
	hoursWorkedByEvent = make(map[uint]HoursWorkedByEvent)

	//Accumulate hours worked sorked by Tag
	var hoursWorkedByTag map[uint]HoursWorkedByTag
	hoursWorkedByTag = make(map[uint]HoursWorkedByTag)

	for _, shift := range volunteerShifts {
		//Get Event info
		var eventInfo types.Event
		db.First(&eventInfo, shift.EventID)
		fmt.Println(eventInfo)

		//Get shift hours
		shiftMinutes := (shift.ActualEndTime - shift.ActualStartTime) / 60000
		shiftHours := shiftMinutes / 60

		//Get NPO info
		var npoInfo types.NPO
		db.Select("npo_name").First(&npoInfo, eventInfo.NPOID)

		//By NPO
		_, ok := hoursWorkedByNPO[eventInfo.NPOID]
		if ok {

			updatedShiftHours := hoursWorkedByNPO[eventInfo.NPOID].HoursWorked + shiftHours
			hoursWorkedByNPO[eventInfo.NPOID] = HoursWorkedByNPO{
				NPOID:       eventInfo.NPOID,
				NPOName:     npoInfo.NPOName,
				HoursWorked: updatedShiftHours,
			}

		} else {
			hoursWorkedByNPO[eventInfo.NPOID] = HoursWorkedByNPO{
				NPOID:       eventInfo.NPOID,
				NPOName:     npoInfo.NPOName,
				HoursWorked: shiftHours,
			}
		}

		//By Event
		_, eventOK := hoursWorkedByEvent[eventInfo.ID]

		if eventOK {
			updatedShiftHours := hoursWorkedByEvent[eventInfo.ID].HoursWorked + shiftHours
			hoursWorkedByEvent[eventInfo.ID] = HoursWorkedByEvent{
				EventID:        eventInfo.ID,
				EventName:      eventInfo.Name,
				EventStartTime: eventInfo.StartTime,
				EventEndTime:   eventInfo.EndTime,
				HoursWorked:    updatedShiftHours,
			}
		} else {
			hoursWorkedByEvent[eventInfo.ID] = HoursWorkedByEvent{
				EventID:        eventInfo.ID,
				EventName:      eventInfo.Name,
				EventStartTime: eventInfo.StartTime,
				EventEndTime:   eventInfo.EndTime,
				HoursWorked:    shiftHours,
			}
		}

		//By Tag

		//Get Tags by event first
		var eventTags []types.Tag
		db.Table("tags").Joins("inner join event_tags on event_tags.tag_id = tags.id").Joins("inner join events on event_tags.event_id = events.id").Where("events.id = ?", eventInfo.ID).Find(&eventTags)
		for _, tag := range eventTags {
			_, tagOK := hoursWorkedByTag[tag.ID]
			if tagOK {
				updatedShiftHours := hoursWorkedByTag[tag.ID].HoursWorked + shiftHours
				hoursWorkedByTag[tag.ID] = HoursWorkedByTag{
					TagID:       tag.ID,
					TagName:     tag.TagName,
					HoursWorked: updatedShiftHours,
				}
			} else {

				hoursWorkedByTag[tag.ID] = HoursWorkedByTag{
					TagID:       tag.ID,
					TagName:     tag.TagName,
					HoursWorked: shiftHours,
				}
			}

		}

	}

	type ReportData struct {
		HoursByNPO   map[uint]HoursWorkedByNPO
		HoursByEvent map[uint]HoursWorkedByEvent
		HoursByTag   map[uint]HoursWorkedByTag
	}

	reportData := ReportData{
		HoursByNPO:   hoursWorkedByNPO,
		HoursByEvent: hoursWorkedByEvent,
		HoursByTag:   hoursWorkedByTag,
	}

	ctx.JSON(reportData)

}
