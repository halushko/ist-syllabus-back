package api

type getAllLevelsResult struct {
	Response
	Result []string `json:"result"`
}
