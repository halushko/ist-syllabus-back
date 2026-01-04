package middleware

import (
	"errors"
	"fmt"
	"ist-syllabus-back/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetAllLevels(db *utils.SQLite) ([]Level, error) {
	res := make([]Level, 0)
	query := `SELECT DISTINCT name, id FROM levels ORDER BY CAST(lvl AS INTEGER) ASC, id;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		l := Level{
			ID:   row["id"].(string),
			Name: row["name"].(string),
		}
		res = append(res, l)
	}
	return res, nil
}

func GetAllStatuses(db *utils.SQLite) ([]IdValue, error) {
	res := make([]IdValue, 0)
	query := `SELECT name, id FROM statuses ORDER BY CAST(id AS INTEGER) ASC;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		l := IdValue{
			ID:    row["id"].(string),
			Value: row["name"].(string),
		}
		res = append(res, l)
	}
	return res, nil
}

func GetAllForms(db *utils.SQLite) ([]IdValue, error) {
	res := make([]IdValue, 0)
	query := `SELECT DISTINCT name, id FROM forms ORDER BY CAST(id AS INTEGER) ASC;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		l := IdValue{
			ID:    row["id"].(string),
			Value: row["name"].(string),
		}
		res = append(res, l)
	}
	return res, nil
}

func GetSemestersCount(level string, db *utils.SQLite) (int, error) {
	query := `SELECT semesters FROM levels WHERE id = ?;`
	rows, err := db.Select(query, level)
	if err != nil {
		log.Errorf("GetSemestersCount err: %v", err)
		return -1, err
	}
	for _, row := range rows {
		if res, err := strconv.Atoi(row["semesters"].(string)); err != nil {
			log.Errorf("GetSemestersCount Atoi err: %v", err)
			return -1, err
		} else {
			return res, nil
		}
	}

	message := fmt.Sprintf("The semesters count for %s is not set", level)
	log.Error(message)
	return -1, errors.New(message)
}
