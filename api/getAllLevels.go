package api

import (
	"ist-syllabus-back/middleware"
	"ist-syllabus-back/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAllStatuses(r *gin.Engine, db *utils.SQLite) {
	r.GET("/api/get-all-levels", func(c *gin.Context) {
		resp := &getAllLevelsResult{}
		res, err := middleware.GetAllLevels(db)
		if err != nil {
			log.Errorf("GetAllLevels err: %v", err)
			resp.AddError(err)
		} else {
			resp.Levels = res
		}

		c.JSON(resp.GetCode(), resp)
	})
}
