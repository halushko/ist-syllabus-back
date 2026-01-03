package api

import (
	"ist-syllabus-back/middleware"
)

type getAllLevelsResult struct {
	Response
	Levels []middleware.Level `json:"levels"`
}

type getAllStatusesResult struct {
	Response
	Statuses []middleware.Status `json:"statuses"`
}
