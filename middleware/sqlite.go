package middleware

import (
	"ist-syllabus-back/utils"

	log "github.com/sirupsen/logrus"
)

func GetAllLevels(db *utils.SQLite) ([]Level, error) {
	res := make([]Level, 0)
	query := `SELECT DISTINCT name, short FROM levels ORDER BY CAST(lvl AS INTEGER) ASC;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		l := Level{
			ID:   row["short"].(string),
			Name: row["name"].(string),
		}
		res = append(res, l)
	}
	return res, nil
}

func GetAllStatuses(db *utils.SQLite) ([]Status, error) {
	res := make([]Status, 0)
	query := `SELECT DISTINCT name, id FROM statuses ORDER BY CAST(id AS INTEGER) ASC;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		l := Status{
			ID:   row["id"].(string),
			Name: row["name"].(string),
		}
		res = append(res, l)
	}
	return res, nil
}
