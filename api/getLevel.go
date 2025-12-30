package api

import (
	"ist-syllabus-back/middleware"
	"ist-syllabus-back/utils"

	"github.com/gin-gonic/gin"
)

func GetAllLevels(r *gin.Engine, db *utils.SQLite) {
	r.GET("/api/get-all-levels", func(c *gin.Context) {
		resp := &getAllLevelsResult{}
		res, err := middleware.GetAllLevels(db)
		if err != nil {
			resp.AddError(err)
		} else {
			resp.Result = res
		}

		c.JSON(resp.GetCode(), resp)
		return
	})
}
