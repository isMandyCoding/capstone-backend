package controllers

import (
	"time"

	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/src/config"
	types "github.com/theycallmethetailor/capstone-backend/src/models"
)

type ReturnEvent struct {
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	NPOID           uint
	NPOName         string
	Name            string
	StartTime       int64
	EndTime         int64
	Tags            []types.Tag
	Description     string
	Location        string
	NumOfVolunteers int
	Shifts          []types.Shift
}

type ReturnNPO struct {
	ID          uint
	NPOName     string
	Description string
	Website     string
	Email       string
	FirstName   string
	LastName    string
	Events      []ReturnEvent
}

func GetAllNPOs(ctx iris.Context) {
	// Create connection to database
	db, _ := databaseConfig.DbStart()
	// Close connection when function is done running
	defer db.Close()

	// Create container for all users
	var npos []types.NPO

	// Query database for all users and populate container with said results
	db.Find(&npos)

	// Respond to request with JSON of all the users
	ctx.JSON(npos)
}

func ShowNPO(ctx iris.Context) {

	// Create connection to database
	db, _ := databaseConfig.DbStart()
	// Close connection when function is done running
	defer db.Close()

	// Create container for one user
	var npo types.NPO

	// Acquire the id via the url params
	urlParam, _ := ctx.Params().GetInt("id")

	// Query database for user with a certain ID
	db.First(&npo, urlParam)

	if npo.ID == 0 {
		ctx.Values().Set("message", "Unable to locate NPO with the provided ID.")
		ctx.StatusCode(500)
	}

	returnNPO := ReturnNPO{
		ID:          npo.ID,
		NPOName:     npo.NPOName,
		Description: npo.Description,
		Website:     npo.Website,
		Email:       npo.Email,
		FirstName:   npo.FirstName,
		LastName:    npo.LastName,
	}

	db.Model(&npo).Order("events.start_time asc").Related(&npo.Events)

	var returnEvents []ReturnEvent

	for _, event := range npo.Events {

		var tags []types.Tag

		db.Table("tags").Joins("inner join event_tags on event_tags.tag_id = tags.id").Joins("inner join events on event_tags.event_id = events.id").Where("events.id = ?", event.ID).Find(&tags)

		var npoInfo types.NPO
		db.Select("npo_name").First(&npoInfo, event.NPOID)

		returnEvent := ReturnEvent{
			ID:              event.ID,
			CreatedAt:       event.CreatedAt,
			UpdatedAt:       event.UpdatedAt,
			NPOID:           event.NPOID,
			NPOName:         npoInfo.NPOName,
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

func CreateNPO(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody types.NPO

	ctx.ReadJSON(&requestBody)

	npo := types.NPO{
		NPOName:     requestBody.NPOName,
		Description: requestBody.Description,
		Website:     requestBody.Website,
		Email:       requestBody.Email,
		FirstName:   requestBody.FirstName,
		LastName:    requestBody.LastName,
		Password:    requestBody.Password,
	}

	db.NewRecord(npo)
	db.Create(&npo)

	if db.NewRecord(npo) == false {
		ctx.JSON(npo)
	} else {
		ctx.Values().Set("message", "Error creating new NPO. Please try again.")
		ctx.StatusCode(500)
	}

}

func UpdateNPO(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var npo types.NPO

	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&npo, urlParam)

	var requestBody types.NPO

	ctx.ReadJSON(&requestBody)

	// Update multiple attributes with `struct`, will only update those changed & non blank fields
	db.Model(&npo).Updates(types.NPO{
		NPOName:     requestBody.NPOName,
		Description: requestBody.Description,
		Website:     requestBody.Website,
		Email:       requestBody.Email,
		FirstName:   requestBody.FirstName,
		LastName:    requestBody.LastName,
		Password:    requestBody.Password,
	})

	var updatedNPO types.NPO

	db.First(&updatedNPO, urlParam)

	ctx.JSON(updatedNPO)

}

func DeleteNPO(ctx iris.Context) {

	//connect to database at start, close db at end of func
	db, _ := databaseConfig.DbStart()
	defer db.Close()

	var npo types.NPO
	urlParam, _ := ctx.Params().GetInt("id")

	db.First(&npo, urlParam)

	db.Unscoped().Delete(&npo)
	ctx.JSON(npo)

}
