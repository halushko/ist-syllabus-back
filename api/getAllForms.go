package api

import (
	"ist-syllabus-back/middleware"
	"ist-syllabus-back/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAllForms(r *gin.Engine, db *utils.SQLite) {
	r.GET("/api/get-all-forms", func(c *gin.Context) {
		resp := &getIdValueResult{}
		res, err := middleware.GetAllForms(db)
		if err != nil {
			log.Errorf("GetAllForms err: %v", err)
			resp.AddError(err)
		} else {
			resp.Values = res
		}

		c.JSON(resp.GetCode(), resp)
	})
}
