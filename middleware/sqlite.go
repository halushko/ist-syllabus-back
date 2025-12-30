package middleware

import (
	"ist-syllabus-back/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetAllLevels(db *utils.SQLite) ([]Level, error) {
	levels := make([]Level, 0)
	query := `SELECT name, lvl FROM levels ORDER BY lvl ASC;`
	rows, err := db.Select(query)
	if err != nil {
		log.Errorf("GetAllLevels err: %v", err)
		return nil, err
	}
	for _, row := range rows {
		lvl, err := strconv.Atoi(row["lvl"].(string))
		if err != nil {
			log.Errorf("GetAllLevels Atoi err for Level column: %v", err)
			return nil, err
		}
		level := Level{
			Name:  row["name"].(string),
			Level: lvl,
		}
		levels = append(levels, level)
	}
	return levels, nil
}
