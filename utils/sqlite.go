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

	res.initTable(googleSheetsConfig.Levels[0], googleSheetsConfig.Levels[1])

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
	err := s.db.DropTable(tableName)
	if err != nil {
		log.Errorf("Can't drop %s sheet: %v", tableName, err)
		panic(err)
	}

	var sb strings.Builder

	i := 0
	for _, col := range headers {
		i++
		if i < len(headers)-1 {
			sb.WriteString("\t" + col + " TEXT,\n")
		} else {
			sb.WriteString("\t" + col)
		}
	}

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s);\n`, tableName, sb.String())
	log.Debugf("Creating table with query:\n%s", query)
	err = s.db.Execute(query)

	if err != nil {
		log.Errorf("Can't create %s table: %v", tableName, err)
		panic(err)
	}
}

func (s *SQLite) fillTable(tableName string, table map[string][]string) {
	var sbc strings.Builder
	var sbv strings.Builder

	i := 0
	j := 0
	for col := range table {
		i++
		if i < len(table)-1 {
			sbc.WriteString(col + ", ")
			sbv.WriteString(table[col][j] + ", ")
		} else {
			sbc.WriteString(col)
			sbv.WriteString(table[col][j])
		}
	}

	query := fmt.Sprintf(`INSERT INTO %s (%s)\nVALUES (%s);\n`, tableName, sbc.String(), sbv.String())
	log.Debugf("Inserting into table with query:\n%s", query)
	err := s.db.Execute(query)
	if err != nil {
		log.Errorf("Can't fill %s table: %v", tableName, err)
		panic(err)
	}
}
