package utils

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/halushko/core-go/google"
	log "github.com/sirupsen/logrus"
)

type GoogleReader struct {
	gc            *google.Client
	spreadsheetId string
}

// InitGoogleReader creates and returns a connection to the Google Spreadsheet.
func InitGoogleReader() (*GoogleReader, error) {
	ctx := context.Background()
	spreadSheetId := os.Getenv("SPREAD_SHEET_ID")
	googleCredentialsPath := os.Getenv("GOOGLE_CREDS_JSON")

	g, err := google.NewFromServiceAccountFile(ctx, googleCredentialsPath)
	if err != nil {
		log.Error("Failed to create Google Sheets client: ", err)
		return nil, err
	}

	return &GoogleReader{gc: g, spreadsheetId: spreadSheetId}, nil
}

func (r *GoogleReader) readGoogleTable(sheetName string) (map[string]string, map[string][]string, error) {
	if r == nil || r.gc == nil {
		return nil, nil, errors.New("client is not initialized")
	}
	if strings.TrimSpace(r.spreadsheetId) == "" {
		return nil, nil, errors.New("spreadsheet ID is empty")
	}
	if strings.TrimSpace(sheetName) == "" {
		return nil, nil, errors.New("sheet name is empty")
	}

	maxCol, err := strconv.Atoi(os.Getenv("MAX_GOOGLE_SHEET_COLS"))
	if err != nil {
		maxCol = 100
	}

	maxRows, err := strconv.Atoi(os.Getenv("MAX_GOOGLE_SHEET_ROWS"))
	if err != nil {
		maxRows = 1000
	}

	res, err := r.gc.ReadByIndexes(r.spreadsheetId, sheetName, 1, 1, maxCol, 1)
	if err != nil {
		log.Error("Failed to read headers from Google Sheets: ", err)
		return nil, nil, err
	}

	size := 0
	for i, h := range res[0] {
		if h == "" {
			size = i + 1
			break
		}
	}
	if size == 0 {
		return nil, nil, errors.New("data not found for table " + sheetName)
	}

	log.Debugf("Detected %d columns for sheet %s", size, sheetName)
	columns := res[0][0:size]

	res, err = r.gc.ReadByIndexes(r.spreadsheetId, sheetName, 1, 2, size, maxRows+1)
	if err != nil {
		log.Error("Failed to read data from Google Sheets: ", err)
		return nil, nil, err
	}

	headers := res[0][0:size]

	headersMap := make(map[string]string)
	table := make(map[string][]string)

	for i := range columns {
		headersMap[headers[i]] = columns[i]
	}

	cleanTable := removeEmptyRows(res[1:])

	for i, h := range columns {
		c := make([]string, 0)
		for _, r := range cleanTable {
			c = append(c, r[i])
		}
		table[h] = c
	}

	return headersMap, table, nil
}

func removeEmptyRows(data [][]string) [][]string {
	cleanedData := make([][]string, 0)
	for _, row := range data {
		if !ifRowIsEmpty(row) {
			cleanedData = append(cleanedData, row)
		}
	}
	return cleanedData
}
func ifRowIsEmpty(row []string) bool {
	for _, v := range row {
		if v != "" {
			return false
		}
	}
	return true
}
