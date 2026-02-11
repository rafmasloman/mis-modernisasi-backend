package model

import "gorm.io/gorm"

type FilterBuilderModel struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Title    string `gorm:"column:title"`
	Query    string `gorm:"column:query"`
	ReportId int    `gorm:"column:report_id"`
	Type     string `gorm:"column:type"`
	Value    string `gorm:"column:value"`
	Label    string `gorm:"column:label"`
	Api      string `gorm:"column:api"`
	ReqBody  string `gorm:"column:req_body"`
	Required bool   `gorm:"column:required"`
	OrderNum int    `gorm:"column:order_num"`
	OnLoad   bool   `gorm:"column:on_load"`
	OnChange string `gorm:"column:on_change"`

	Report ReportBuilderModel `gorm:"foreignKey:ReportId;references:ReportId"`
}

func (FilterBuilderModel) TableName() string {
	return "filter_builder"
}
