package entity

type LookupBranchFlag struct {
	KonvenCd string
	KonvenNm string
}

type LookupBranchSubBranch struct {
	SubBranchCd string
	SubBranchNm string
	Flag        string
}

type LookupBranchs struct {
	BranchCd string
	BranchNm string
	Flag     string
}

type LookupBranchArea struct {
	AreaCd string
	AreaNm string
	Flag   string
}

type LookupBranchRegion struct {
	RegionCd string
	RegionNm string
	Flag     string
}
