package models

import "time"

type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float64
	Data       map[string]interface{}
	Title      string
	Heading    string
	Subheading string
	Year       time.Time
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
}
