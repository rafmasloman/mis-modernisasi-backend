package usecase

import (
	"fmt"
	"strconv"

	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/dto"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/entity"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/repositories"
)

type FilterBuilderUsecaseImpl struct {
	repo       repositories.FilterBuilderRepository
	reportRepo repositories.ReportBuilderRepository
}

type FilterBuilderUsecase interface {
	CreateFilterBuilder(dto dto.FilterBuilderDTO) error
	GetAllFilterBuilder() (*[]dto.FilterBuilderResponseDTO, error)
	GetFilterBuilderByReportId(reportId string) (*[]dto.FilterBuilderResponseDTO, error)
	DeleteFilterBuilder(id string) error
	UpdateFilterBuilder(id string, dto dto.FilterBuilderDTO) error
}

func NewFilterBuilderUsecase(repo repositories.FilterBuilderRepository, reportRepo repositories.ReportBuilderRepository) FilterBuilderUsecase {
	return &FilterBuilderUsecaseImpl{repo: repo, reportRepo: reportRepo}
}

func (u *FilterBuilderUsecaseImpl) GetAllFilterBuilder() (*[]dto.FilterBuilderResponseDTO, error) {
	data, err := u.repo.GetAllFilterBuilder()

	if err != nil {
		return nil, err
	}

	entities := make([]dto.FilterBuilderResponseDTO, 0, len(*data))

	for _, item := range *data {
		entities = append(entities, dto.FilterBuilderResponseDTO{
			Id:       item.ID,
			ReportId: item.ReportId,
			Name:     item.Name,
			Query:    item.Query,
			Title:    item.Title,
			Type:     item.Type,
			Required: item.Required,
			OrderNum: item.OrderNum,
		})
	}

	return &entities, nil

}

func (u *FilterBuilderUsecaseImpl) CreateFilterBuilder(dto dto.FilterBuilderDTO) error {

	filterData := entity.FilterBuilder{
		Name:     dto.Name,
		Title:    dto.Title,
		Query:    dto.Query,
		ReportId: dto.ReportId,
		Type:     dto.Type,
		Required: dto.Required,
		OrderNum: dto.OrderNum,
	}

	_, err := u.reportRepo.FindReportBuilderById(dto.ReportId)

	if err != nil {
		return err
	}

	if err := u.repo.CreateFilterBuilder(filterData); err != nil {
		return fmt.Errorf(`failed to create filter : %v`, err)
	}

	return nil
}

func (u *FilterBuilderUsecaseImpl) GetFilterBuilderByReportId(reportId string) (*[]dto.FilterBuilderResponseDTO, error) {

	convertReportId, err := strconv.Atoi(reportId)

	if err != nil {
		return nil, fmt.Errorf(`failed to convert report id : %v `, err)
	}

	filterResult, err := u.repo.GetFilterBuilderByReportId(convertReportId)

	if err != nil {
		return nil, err
	}

	entities := make([]dto.FilterBuilderResponseDTO, 0, len(*filterResult))

	for _, item := range *filterResult {
		entities = append(entities, dto.FilterBuilderResponseDTO{
			Id:       item.ID,
			ReportId: item.ReportId,
			Name:     item.Name,
			Query:    item.Query,
			Title:    item.Title,
			Type:     item.Type,
			Required: item.Required,
			OrderNum: item.OrderNum,
		})
	}

	return &entities, nil

}

func (u *FilterBuilderUsecaseImpl) DeleteFilterBuilder(id string) error {

	convertId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf(`failed to convert id : %v `, err)
	}

	if err := u.repo.DeleteFilterBuilder(convertId); err != nil {
		return fmt.Errorf(`failed to delete filter builder %v `, err)
	}

	return nil
}

func (u *FilterBuilderUsecaseImpl) UpdateFilterBuilder(id string, dto dto.FilterBuilderDTO) error {

	convertId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf(`failed to convert id : %v `, err)
	}

	filterUpdate := entity.FilterBuilder{
		Name:     dto.Name,
		Title:    dto.Title,
		Query:    dto.Query,
		Type:     dto.Type,
		ReportId: dto.ReportId,
		Required: dto.Required,
		OrderNum: dto.OrderNum,
	}

	if err := u.repo.UpdateFilterBuilder(convertId, filterUpdate); err != nil {
		return fmt.Errorf(`failed to update filter : %v`, err)
	}

	return nil
}
