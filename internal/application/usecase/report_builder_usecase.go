package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/dto"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/entity"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/repositories"
)

type ReportBuilderUsecaseImpl struct {
	repo repositories.ReportBuilderRepository
}

type ReportBuilderUsecase interface {
	CreateReportBuilder(dto dto.DTOCreateReportBuilder) error
	GetReportByRouterName(routerName string) (*dto.DTOReportBuilderResponse, error)
	DeleteReport(reportId string) error
	UpdateReport(reportId string, dto dto.DTOCreateReportBuilder) error
	GetAllReportBuilder() (*[]dto.DTOReportBuilderResponse, error)
}

func NewReportBuilderUsecase(repo repositories.ReportBuilderRepository) ReportBuilderUsecase {
	return &ReportBuilderUsecaseImpl{repo: repo}
}

func (u *ReportBuilderUsecaseImpl) CreateReportBuilder(dto dto.DTOCreateReportBuilder) error {

	q := strings.ToLower(dto.Query)

	// if dto.Query == "" {
	// 	return fmt.Errorf(`Query cannot be empty`)
	// }

	if !strings.HasPrefix(strings.TrimSpace(q), "select") {
		return fmt.Errorf(`only SELECT query allowed`)
	}

	reportData := entity.ReportBuilder{
		Name:        dto.Name,
		Query:       dto.Query,
		RouteName:   dto.RouteName,
		Description: dto.Description,
		Columns:     dto.Columns,
		IsActive:    dto.IsActive,
	}

	err := u.repo.CreateReportBuilder(reportData)

	if err != nil {
		return err
	}

	return nil

}

func (u *ReportBuilderUsecaseImpl) GetAllReportBuilder() (*[]dto.DTOReportBuilderResponse, error) {

	data, err := u.repo.GetAllReportBuilder()

	if err != nil {
		return nil, fmt.Errorf(`failed to fetch all report builder : %v `, err)
	}

	entities := make([]dto.DTOReportBuilderResponse, 0, len(*data))

	for _, item := range *data {
		entities = append(entities, dto.DTOReportBuilderResponse{
			ReportId:    item.ReportId,
			Name:        item.Name,
			Query:       item.Query,
			RouteName:   item.RouteName,
			Description: item.Description,
			Columns:     item.Columns,
			IsActive:    item.IsActive,
		})
	}

	return &entities, nil

}

func (u *ReportBuilderUsecaseImpl) GetReportByRouterName(routerName string) (*dto.DTOReportBuilderResponse, error) {
	result, err := u.repo.GetReportByRouterName(routerName)

	if err != nil {
		return nil, err
	}

	mappingToResponse := dto.DTOReportBuilderResponse{
		ReportId:    result.ReportId,
		Name:        result.Name,
		Query:       result.Query,
		RouteName:   result.RouteName,
		Description: result.Description,
		Columns:     result.Columns,
		IsActive:    result.IsActive,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	return &mappingToResponse, nil
}

func (u *ReportBuilderUsecaseImpl) DeleteReport(reportId string) error {

	convertReportId, err := strconv.Atoi(reportId)

	if err != nil {
		return fmt.Errorf(`failed to convert: %v`, err)
	}

	if err := u.repo.DeleteReport(convertReportId); err != nil {
		return err
	}

	return nil
}

func (u *ReportBuilderUsecaseImpl) UpdateReport(reportId string, dto dto.DTOCreateReportBuilder) error {

	convertReportId, err := strconv.Atoi(reportId)

	if err != nil {
		return fmt.Errorf(`failed to convert: %v`, err)
	}

	payloadUpdate := entity.ReportBuilder{
		Name:        dto.Name,
		Query:       dto.Query,
		RouteName:   dto.RouteName,
		Description: dto.Description,
		Columns:     dto.Columns,
		IsActive:    dto.IsActive,
	}

	if err := u.repo.UpdateReport(convertReportId, payloadUpdate); err != nil {
		return err
	}

	return nil

}
