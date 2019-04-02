package controllers

import (
	"github.com/kataras/iris"
	databaseConfig "github.com/theycallmethetailor/capstone-backend/src/config"
	types "github.com/theycallmethetailor/capstone-backend/src/models"
)

func GetAllTags(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var tags []types.Tag

	db.Find(&tags)

	ctx.JSON(tags)
}

func GetTagList(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var tags []types.Tag

	var requestBody []uint

	ctx.ReadJSON(&requestBody)

	for _, tagID := range requestBody {
		var tag types.Tag
		db.First(&tag, tagID)

		tags = append(tags, tag)
	}

	ctx.JSON(tags)
}

func CreateTags(ctx iris.Context) {
	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody []string

	ctx.ReadJSON(&requestBody)

	var returnTags []types.Tag

	for _, newTag := range requestBody {

		tag := types.Tag{
			TagName: newTag,
		}
		//Only create a new tag if the tag doesn't already exist
		db.FirstOrCreate(&tag, types.Tag{TagName: newTag})

		returnTags = append(returnTags, tag)
	}

	ctx.JSON(returnTags)
}

func DeleteTags(ctx iris.Context) {

	db, _ := databaseConfig.DbStart()

	defer db.Close()

	var requestBody []uint

	ctx.ReadJSON(&requestBody)

	for _, tagID := range requestBody {
		var tag types.Tag
		db.First(&tag, tagID)
		db.Unscoped().Delete(&tag)
	}

	ctx.JSON(requestBody)

}
