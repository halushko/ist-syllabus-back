package middleware

import (
	"ist-syllabus-back/utils"

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
		level := Level{
			Name:  row["name"].(string),
			Level: int(row["lvl"].(int64)),
		}
		levels = append(levels, level)
	}
	return levels, nil
}
