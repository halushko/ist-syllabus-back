package api

import (
	"ist-syllabus-back/middleware"
)

type getAllLevelsResult struct {
	Response
	Levels []middleware.Level `json:"levels"`
}

type getIdValueResult struct {
	Response
	Values []middleware.IdValue `json:"values"`
}
