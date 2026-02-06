package model

import (
	"time"

	"gorm.io/datatypes"
)

type ReportBuilderModel struct {
	ReportId    int            `gorm:"column:report_id;primaryKey;autoIncrement"`
	Name        string         `gorm:"column:name"`
	Query       string         `gorm:"column:query"`
	RouteName   string         `gorm:"column:route_name"`
	Description string         `gorm:"column:description"`
	Columns     datatypes.JSON `gorm:"column:columns;type:jsonb"`
	IsActive    bool           `gorm:"column:is_active"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
}

func (ReportBuilderModel) TableName() string {
	return "report_builder"
}
