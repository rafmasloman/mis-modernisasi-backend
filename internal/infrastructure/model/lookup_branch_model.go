package model

import "gorm.io/gorm"

type LookupBranchModel struct {
	gorm.Model
	SubBranchCd string `gorm:"column:sub_branch_cd"`
	SubBranchNm string `gorm:"column:sub_branch_nm"`
	BranchCd    string `gorm:"column:branch_cd"`
	BranchNm    string `gorm:"column:branch_nm"`
	AreaCd      string `gorm:"column:area_cd"`
	AreaNm      string `gorm:"column:area_nm"`
	RegionCd    string `gorm:"column:region_cd"`
	RegionNm    string `gorm:"column:region_nm"`
	Flag        string `gorm:"column:flag"`
}

func (LookupBranchModel) TableName() string {
	return "lookup_branch"
}
