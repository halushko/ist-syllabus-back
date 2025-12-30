package middleware

import (
	"ist-syllabus-back/utils"

	log "github.com/sirupsen/logrus"
)

func GetAllLevels(db *utils.SQLite) ([]string, error) {
	res := make([]string, 0)
	query := `SELECT DISTINCT name FROM levels ORDER BY CAST(lvl AS INTEGER) ASC;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		res = append(res, row["name"].(string))
	}
	return res, nil
}
