package usecase

import (
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/dto"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/repositories"
)

type LookupBranchUsecaseImpl struct {
	repo repositories.LookupBranchRepository
}

type LookupBranchUsecase interface {
	GetKonven() (*[]dto.LookupBranchFlagResponseDTO, error)
	GetWilayah(flag string) (*[]dto.LookupBranchRegionResponseDTO, error)
}

func NewLookupBranchUsecase(repo repositories.LookupBranchRepository) LookupBranchUsecase {
	return &LookupBranchUsecaseImpl{repo: repo}
}

func (u *LookupBranchUsecaseImpl) GetKonven() (*[]dto.LookupBranchFlagResponseDTO, error) {
	result, err := u.repo.GetKonven()

	if err != nil {
		return nil, err
	}

	entities := make([]dto.LookupBranchFlagResponseDTO, 0, len(*result))

	for _, item := range *result {
		entities = append(entities, dto.LookupBranchFlagResponseDTO{
			KonvenCd: item.KonvenCd,
			KonvenNm: item.KonvenNm,
		})
	}

	return &entities, nil
}

func (u *LookupBranchUsecaseImpl) GetWilayah(flag string) (*[]dto.LookupBranchRegionResponseDTO, error) {

	result, err := u.repo.GetWilayah(flag)

	if err != nil {
		return nil, err
	}

	entities := make([]dto.LookupBranchRegionResponseDTO, 0, len(*result))

	for _, item := range *result {
		entities = append(entities, dto.LookupBranchRegionResponseDTO{
			RegionCd: item.RegionCd,
			RegionNm: item.RegionNm,
		})
	}

	return &entities, nil
}
