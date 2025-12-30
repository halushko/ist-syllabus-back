package main

import (
	"errors"
	"ist-syllabus-back/api"
	"ist-syllabus-back/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/halushko/core-go/logger"
	log "github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

func main() {
	err := logger.Init()
	if err != nil {
		panic(err)
	}

	db, err := utils.InitSQLite()
	if err != nil {
		log.Errorf("Failed to initialize database: %v", err)
		panic(err)
	}

	r := gin.Default()
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       15 * time.Minute,
		WriteTimeout:      15 * time.Minute,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 1 * time.Minute,
	}
	api.GetAllLevels(r, db)

	log.Infof("Starting server with long timeouts…")

	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Errorf("[FATAL] %v", err)
		panic(err)
	}
	//c, err := google.NewFromServiceAccountFile(ctx, GoogleCredentialsPath)
	//if err != nil {
	//	log.Error("Failed to create Google Sheets client: ", err)
	//	return
	//}
	//res, err := c.ReadByA1(ctx, SpreadSheetId, "Кафедри", "A", 1, "B", 4)
	//if err != nil {
	//	log.Error("Failed to read data from Google Sheets: ", err)
	//	return
	//}
	//log.Infof("Read data: %v", res)
}
