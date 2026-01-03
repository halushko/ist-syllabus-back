package middleware

type Level struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Level int    `json:"level,omitempty"`
}

type Status struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}
