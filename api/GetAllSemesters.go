package api

import (
	"ist-syllabus-back/middleware"
	"ist-syllabus-back/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAllSemesters(r *gin.Engine, db *utils.SQLite) {
	r.GET("/api/get-all-semesters", func(c *gin.Context) {
		level := c.Query("level")

		resp := &getIdValueResult{}
		cnt, err := middleware.GetSemestersCount(level, db)
		if err != nil {
			log.Errorf("GetAllSemesters err: %v", err)
			resp.AddError(err)
		} else {
			resp.Values = make([]middleware.IdValue, cnt)
			for i := 0; i < cnt; i++ {
				resp.Values[i] = middleware.IdValue{
					ID:    strconv.Itoa(i),
					Value: strconv.Itoa(i + 1),
				}
			}
		}

		c.JSON(resp.GetCode(), resp)
	})
}
