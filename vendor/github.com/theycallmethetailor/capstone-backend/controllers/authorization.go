package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

// type SafeNPO struct {
// 	CreatedAt time.Time
// 	UpdatedAt time.Time

// }

func Login(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody struct {
		UserType string
		Email    string
		Password string
	}

	ctx.ReadJSON(&requestBody)

	if requestBody.UserType == "NPO" {
		var foundNPO types.NPO
		fmt.Print("requestBody: ", requestBody)
		db.Where("email = ? AND password = ?", requestBody.Email, requestBody.Password).First(&foundNPO)
		if db.Where("email = ? AND password = ?", requestBody.Email, requestBody.Password).First(&foundNPO).RecordNotFound() {

			ctx.Values().Set("message", "Invalid credentials/user not found.")
			ctx.StatusCode(500)
		} else {
			db.Model(&foundNPO).Order("events.start_time asc").Related(&foundNPO.Events)
			returnNPO := ReturnNPO{
				ID:          foundNPO.ID,
				NPOName:     foundNPO.NPOName,
				Description: foundNPO.Description,
				Website:     foundNPO.Website,
				Email:       foundNPO.Email,
				FirstName:   foundNPO.FirstName,
				LastName:    foundNPO.LastName,
			}

			var returnEvents []ReturnEvent

			for _, event := range foundNPO.Events {

				var tags []types.Tag

				db.Table("tags").Joins("inner join event_tags on event_tags.tag_id = tags.id").Joins("inner join events on event_tags.event_id = events.id").Where("events.id = ?", event.ID).Find(&tags)

				returnEvent := ReturnEvent{
					ID:              event.ID,
					CreatedAt:       event.CreatedAt,
					UpdatedAt:       event.UpdatedAt,
					NPOID:           event.NPOID,
					NPOName:         returnNPO.NPOName,
					Name:            event.Name,
					StartTime:       event.StartTime,
					EndTime:         event.EndTime,
					Description:     event.Description,
					Location:        event.Location,
					NumOfVolunteers: event.NumOfVolunteers,
					Tags:            tags,
				}

				returnEvents = append(returnEvents, returnEvent)
			}

			returnNPO.Events = returnEvents

			ctx.JSON(returnNPO)
		}
	} else if requestBody.UserType == "Volunteer" {
		var foundVolunteer types.Volunteer

		db.Where("email = ? AND password = ?", requestBody.Email, requestBody.Password).First(&foundVolunteer)

		if db.Where("email = ? AND password = ?", requestBody.Email, requestBody.Password).First(&foundVolunteer).RecordNotFound() {
			fmt.Println("Volunteer not found.")
			ctx.Values().Set("Message", "Invalid credentials/user not found.")
			ctx.StatusCode(500)
		} else {

			db.Model(&foundVolunteer).Related(&foundVolunteer.Shifts)

			returnVolunteer := ReturnVolunteer{
				ID:        foundVolunteer.ID,
				Username:  foundVolunteer.Username,
				Bio:       foundVolunteer.Bio,
				Email:     foundVolunteer.Email,
				FirstName: foundVolunteer.FirstName,
				LastName:  foundVolunteer.LastName,
			}

			for _, shift := range foundVolunteer.Shifts {
				var eventInfo types.Event
				db.First(&eventInfo, shift.EventID)
				var npoInfo types.NPO
				db.First(&npoInfo, eventInfo.NPOID)
				duration := (shift.ActualEndTime - shift.ActualStartTime) / 60000
				returnShift := ReturnShift{
					ID:               shift.ID,
					CreatedAt:        shift.CreatedAt,
					UpdatedAt:        shift.UpdatedAt,
					VolunteerID:      shift.VolunteerID,
					EventID:          shift.EventID,
					WasWorked:        shift.WasWorked,
					ActualStartTime:  shift.ActualStartTime,
					ActualEndTime:    shift.ActualEndTime,
					NPOName:          npoInfo.NPOName,
					EventName:        eventInfo.Name,
					EventDescription: eventInfo.Description,
					EventLocation:    eventInfo.Location,
					NumOfVolunteers:  eventInfo.NumOfVolunteers,
					Duration:         duration,
				}
				returnVolunteer.Shifts = append(returnVolunteer.Shifts, returnShift)
			}

			ctx.JSON(returnVolunteer)

		}

	} else {
		ctx.Values().Set("Message", "Invalid credentials/user not found.")
		ctx.StatusCode(500)
	}
}
