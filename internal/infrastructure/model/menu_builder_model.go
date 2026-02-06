package model

import "gorm.io/gorm"

type MenuBuilderModel struct {
	gorm.Model
	Level           string `gorm:"column:level;type:char(2)"`
	MenuName        string `gorm:"column:menu_name"`
	Parent          string `gorm:"column:parent;type:char(50)"`
	Group           string `gorm:"column:group"`
	Path            string `gorm:"column:path"`
	RouteName       string `gorm:"column:route_name"`
	Component       string `gorm:"column:component"`
	Icon            string `gorm:"column:icon"`
	IsRatingEnabled bool   `gorm:"column:is_rating_enabled;type:tinyint"`
}

func (MenuBuilderModel) TableName() string {
	return "menu_builder"
}
