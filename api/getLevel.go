package api

import (
	"ist-syllabus-back/utils"

	"github.com/gin-gonic/gin"
)

func GetAllLevels(r *gin.Engine, db *utils.SQLite) {
	r.GET("/api/get-all-levels", func(c *gin.Context) {
		//db.
		//g, err := google.NewFromServiceAccountFile(googleCredentialsPath)
		//if err != nil {
		//	log.Error("Failed to create Google Sheets client: ", err)
		//	c.JSON(http.StatusBadRequest, nil)
		//	return
		//}
		//res, err := g.ReadByA1(ctx, spreadSheetId, sheets.Levels, "A", 2, "B", 4)
		//if err != nil {
		//	log.Error("Failed to read data from Google Sheets: ", err)
		//	return
		//}

	})
}
