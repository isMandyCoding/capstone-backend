package controllers

import (
	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/config"
	types "github.com/theycallmethetailor/capstone-backend/models"
)

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
