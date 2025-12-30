package utils

import (
	"fmt"
	"strings"

	"github.com/halushko/core-go/sqlite"
	_ "modernc.org/sqlite"

	log "github.com/sirupsen/logrus"
)

type SQLite struct {
	db   *sqlite.Client
	g    *GoogleReader
	g2db map[string]map[string]string
	db2g map[string]map[string]string
}

// InitSQLite creates and returns a connection to the SQLite database.
//
//goland:noinspection GoResourceLeak
func InitSQLite() (*SQLite, error) {
	sql, err := sqlite.Open(specialityDbName)
	if err != nil {
		return nil, err
	}

	gr, err := InitGoogleReader()
	if err != nil {
		return nil, err
	}

	res := &SQLite{db: sql, g: gr, g2db: make(map[string]map[string]string), db2g: make(map[string]map[string]string)}

	res.initTable(GoogleSheetsConfig.Levels[0], GoogleSheetsConfig.Levels[1])
	res.initTable(GoogleSheetsConfig.Departments[0], GoogleSheetsConfig.Departments[1])
	res.initTable(GoogleSheetsConfig.Specialties[0], GoogleSheetsConfig.Specialties[1])
	res.initTable(GoogleSheetsConfig.Programs[0], GoogleSheetsConfig.Programs[1])
	res.initTable(GoogleSheetsConfig.DepartmentPrograms[0], GoogleSheetsConfig.DepartmentPrograms[1])

	return res, nil
}

func (s *SQLite) initTable(sheetName, tableName string) {
	headers, table, err := s.g.readGoogleTable(sheetName)
	if err != nil {
		log.Errorf("Can't read sheet %s: %v", sheetName, err)
		panic(err)
	}

	gHeaders := make(map[string]string)
	for k, v := range headers {
		gHeaders[v] = k
	}
	s.g2db[sheetName] = headers
	s.db2g[tableName] = gHeaders

	s.recreateTable(tableName, headers)
	s.fillTable(tableName, table)

	log.Infof("Initialized table %s from sheet %s", tableName, sheetName)
}

func (s *SQLite) recreateTable(tableName string, headers map[string]string) {
	log.Debugf("Heading to create table %s: %v", tableName, headers)
	err := s.db.DropTable(tableName)
	if err != nil {
		log.Errorf("Can't drop %s sheet: %v", tableName, err)
		panic(err)
	}

	i := 0
	var sb strings.Builder

	for _, col := range headers {
		i++
		sb.WriteString(col + " TEXT")

		if i < len(headers) {
			sb.WriteString(", ")
		}
	}

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s);`, tableName, sb.String())
	log.Debugf("Creating table with query:\n%s", query)
	err = s.db.Execute(query)

	if err != nil {
		log.Errorf("Can't create %s table: %v", tableName, err)
		panic(err)
	}
}

func (s *SQLite) fillTable(tableName string, table map[string][]string) {
	size := 0

	for row := range table {
		size = len(table[row])
		break
	}

	for j := 0; j < size; j++ {
		var sbc strings.Builder
		var sbv strings.Builder
		i := 0

		for col := range table {
			i++

			sbc.WriteString(col)
			sbv.WriteString("'" + table[col][j] + "'")

			if i < len(table) {
				sbc.WriteString(", ")
				sbv.WriteString(", ")
			}
		}

		query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s);`, tableName, sbc.String(), sbv.String())
		log.Debugf("Inserting into table with query:\n%s", query)
		err := s.db.Execute(query)
		if err != nil {
			log.Errorf("Can't fill %s table: %v", tableName, err)
			panic(err)
		}
	}
}

func (s *SQLite) Select(query string, args ...any) ([]map[string]any, error) {
	return s.db.ExecSelect(query, args...)
}
