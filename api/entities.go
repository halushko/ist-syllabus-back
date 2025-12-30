package api

import (
	"ist-syllabus-back/middleware"
)

type getAllLevelsResult struct {
	Response
	Result []middleware.Level `json:"result"`
}
