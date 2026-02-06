package entity

import (
	"time"

	"gorm.io/datatypes"
)

type ReportBuilder struct {
	ReportId    int
	Name        string
	Query       string
	RouteName   string
	Description string
	Columns     datatypes.JSON
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
