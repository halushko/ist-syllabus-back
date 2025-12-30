package middleware

import "ist-syllabus-back/utils"

func GetAllLevels(db *utils.SQLite) ([]Level, error) {
	levels := make([]Level, 0)
	query := `SELECT name, lvl FROM levels ORDER BY level ASC;`
	rows, err := db.Select(query)
	if err != nil {
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
