package model

import "gorm.io/gorm"

type FilterBuilderModel struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Title    string `gorm:"column:title"`
	Query    string `gorm:"column:query"`
	ReportId int    `gorm:"column:report_id"`
	Type     string `gorm:"column:type"`
	Required bool   `gorm:"column:required"`
	OrderNum int    `gorm:"column:order_num"`

	Report ReportBuilderModel `gorm:"foreignKey:ReportId;references:ReportId"`
}

func (FilterBuilderModel) TableName() string {
	return "filter_builder"
}
