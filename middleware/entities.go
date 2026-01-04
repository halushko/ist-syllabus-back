package middleware

type Level struct {
	ID    string `json:"id,omitempty"`
	Short string `json:"short,omitempty"`
	Name  string `json:"name"`
	Level int    `json:"level,omitempty"`
}

type IdValue struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value"`
}
