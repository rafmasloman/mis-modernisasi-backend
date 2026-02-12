package repositories

import (
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/entity"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/model"
	"gorm.io/gorm"
)

type LookupBranchRepositoryImpl struct {
	db *gorm.DB
}

type LookupBranchRepository interface {
	GetKonven() (*[]entity.LookupBranchFlag, error)
	GetWilayah(flag string) (*[]entity.LookupBranchRegion, error)
}

func NewLookupBranchRepository(db *gorm.DB) LookupBranchRepository {
	return &LookupBranchRepositoryImpl{db: db}
}

func (r *LookupBranchRepositoryImpl) GetKonven() (*[]entity.LookupBranchFlag, error) {
	var model []model.LookupBranchModel

	if err := r.db.Select(`flag as konven_nm, flag as konven_cd`).
		Find(&model).
		Error; err != nil {
		return nil, err
	}

	entities := make([]entity.LookupBranchFlag, 0, len(model))

	for _, item := range model {
		entities = append(entities, entity.LookupBranchFlag{
			KonvenCd: item.Flag,
			KonvenNm: item.Flag,
		})
	}

	return &entities, nil
}

func (r *LookupBranchRepositoryImpl) GetWilayah(flag string) (*[]entity.LookupBranchRegion, error) {
	var model []model.LookupBranchModel

	query := r.db.Select(`region_cd, region_nm`)

	if flag != "" {
		query = query.Where(`flag = ?`, flag)
	}

	if err := query.Find(&model).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.LookupBranchRegion, 0, len(model))

	for _, item := range model {
		entities = append(entities, entity.LookupBranchRegion{
			RegionCd: item.RegionCd,
			RegionNm: item.RegionNm,
		})
	}

	return &entities, nil

}
