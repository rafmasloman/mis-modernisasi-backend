package dto

import (
	"time"

	"gorm.io/datatypes"
)

type DTOCreateReportBuilder struct {
	Name        string         `json:"name" binding:"required"`
	Query       string         `json:"query" binding:"required"`
	RouteName   string         `json:"route_name" binding:"required"`
	Description string         `json:"description"`
	Columns     datatypes.JSON `json:"columns"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type DTOReportBuilderResponse struct {
	ReportId    int            `json:"report_id"`
	Name        string         `json:"name"`
	Query       string         `json:"query"`
	RouteName   string         `json:"route_name"`
	Description string         `json:"description"`
	Columns     datatypes.JSON `json:"columns"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
