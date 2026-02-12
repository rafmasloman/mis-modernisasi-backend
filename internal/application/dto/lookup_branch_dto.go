package dto

type LookupBranchFlagResponseDTO struct {
	KonvenCd string `json:"konven_cd"`
	KonvenNm string `json:"konven_nm"`
}

type LookupBranchSubBranchResponseDTO struct {
	SubBranchCd string `json:"sub_branch_cd"`
	SubBranchNm string `json:"sub_branch_nm"`
	Flag        string `json:"flag"`
}

type LookupBranchsResponseDTO struct {
	BranchCd string `json:"branch_cd"`
	BranchNm string `json:"branch_nm"`
	Flag     string `json:"flag"`
}

type LookupBranchAreaResponseDTO struct {
	AreaCd string `json:"area_cd"`
	AreaNm string `json:"area_nm"`
	Flag   string `json:"flag"`
}

type LookupBranchRegionResponseDTO struct {
	RegionCd string `json:"region_cd"`
	RegionNm string `json:"region_nm"`
	Flag     string `json:"flag"`
}
