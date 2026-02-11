package entity

import "gorm.io/gorm"

type FilterBuilder struct {
	gorm.Model
	Name     string
	Title    string
	Query    string
	ReportId int
	Type     string
	Value    string
	Label    string
	Api      string
	Required bool
	OrderNum int
	OnLoad   bool
	OnChange string
}
