package main

import (
	"context"
	"os"

	"github.com/halushko/core-go/google"
	"github.com/halushko/core-go/logger"
	log "github.com/sirupsen/logrus"
)

var SpreadSheetId = os.Getenv("SPREAD_SHEET_ID")
var GoogleCredentialsPath = os.Getenv("GOOGLE_CREDS_JSON")

func main() {
	err := logger.Init()
	if err != nil {
		log.Error("Failed to create Google Sheets client: ", err)
		panic(err)
	}

	ctx := context.Background()
	c, err := google.NewFromServiceAccountFile(ctx, GoogleCredentialsPath)
	if err != nil {
		log.Error("Failed to create Google Sheets client: ", err)
		return
	}
	res, err := c.ReadByA1(ctx, SpreadSheetId, "Кафедри", "A", 1, "B", 4)
	if err != nil {
		log.Error("Failed to read data from Google Sheets: ", err)
		return
	}
	log.Infof("Read data: %v", res)
}
