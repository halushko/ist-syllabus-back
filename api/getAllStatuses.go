package api

import (
	"ist-syllabus-back/middleware"
	"ist-syllabus-back/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAllStatuses(r *gin.Engine, db *utils.SQLite) {
	r.GET("/api/get-all-statuses", func(c *gin.Context) {
		resp := &getIdValueResult{}
		res, err := middleware.GetAllStatuses(db)
		if err != nil {
			log.Errorf("GetAllStatuses err: %v", err)
			resp.AddError(err)
		} else {
			resp.Values = res
		}

		c.JSON(resp.GetCode(), resp)
	})
}
