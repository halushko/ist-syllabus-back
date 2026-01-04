package main

import (
	"errors"
	"ist-syllabus-back/api"
	"ist-syllabus-back/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:5173",
			"http://localhost:5173",
			"http://127.0.0.1",
			"http://localhost",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: false,
	}))

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       15 * time.Minute,
		WriteTimeout:      15 * time.Minute,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 1 * time.Minute,
	}

	api.GetAllLevels(r, db)
	api.GetAllStatuses(r, db)
	api.GetAllForms(r, db)
	api.GetAllSemesters(r, db)

	log.Infof("Starting server with long timeouts…")

	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Errorf("[FATAL] %v", err)
		panic(err)
	}
}
