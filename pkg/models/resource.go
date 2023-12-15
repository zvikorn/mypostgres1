package models

import "time"

type Resource struct {
	URN          string
	ResourceType string
	Name         string
	Date         time.Time
}
