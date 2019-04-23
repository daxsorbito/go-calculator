package model

import "time"

// Report struct holder for report object
type Report struct {
	Type        string    `json:"type"`
	Message     string    `json:"message"`
	CreatedDate time.Time `json:"createdDate"`
}
