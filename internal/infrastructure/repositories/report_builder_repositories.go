package repositories

import (
	"errors"
	"fmt"

	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/entity"
	domainErrors "github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/errors"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/model"
	"gorm.io/gorm"
)

type ReportBuilderRepositoryImpl struct {
	db *gorm.DB
}

type ReportBuilderRepository interface {
	CreateReportBuilder(item entity.ReportBuilder) error
	GetAllReportBuilder() (*[]entity.ReportBuilder, error)
	GetReportByRouterName(routerName string) (*entity.ReportBuilder, error)
	FindReportBuilderById(reportId int) (*entity.ReportBuilder, error)
	DeleteReport(reportId int) error
	UpdateReport(reportId int, item entity.ReportBuilder) error
}

func NewReportBuilderRepositories(db *gorm.DB) ReportBuilderRepository {
	return &ReportBuilderRepositoryImpl{db: db}
}

func (r *ReportBuilderRepositoryImpl) CreateReportBuilder(item entity.ReportBuilder) error {

	reportData := model.ReportBuilderModel{
		Name:        item.Name,
		Query:       item.Query,
		Description: item.Description,
		RouteName:   item.RouteName,
		Columns:     item.Columns,
		IsActive:    item.IsActive,
	}

	if err := r.db.Create(&reportData).Error; err != nil {
		return err
	}

	return nil
}

func (r *ReportBuilderRepositoryImpl) GetAllReportBuilder() (*[]entity.ReportBuilder, error) {
	var model []model.ReportBuilderModel

	if err := r.db.Find(&model).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.ReportBuilder, 0, len(model))

	for _, item := range model {
		entities = append(entities, entity.ReportBuilder{
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

func (r *ReportBuilderRepositoryImpl) FindReportBuilderById(reportId int) (*entity.ReportBuilder, error) {
	var model model.ReportBuilderModel

	if err := r.db.Table(`report_builder`).Where(`report_id = ?`, reportId).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrReportNotFound
		}
		return nil, err
	}

	entity := entity.ReportBuilder{
		Name:        model.Name,
		Query:       model.Query,
		RouteName:   model.RouteName,
		ReportId:    model.ReportId,
		Description: model.Description,
		Columns:     model.Columns,
		IsActive:    model.IsActive,
		CreatedAt:   model.CreatedAt,
	}

	return &entity, nil
}

func (r *ReportBuilderRepositoryImpl) GetReportByRouterName(routerName string) (*entity.ReportBuilder, error) {
	var model model.ReportBuilderModel

	if err := r.db.Table(`report_builder`).
		Where(`route_name = ?`, routerName).
		First(&model).Error; err != nil {
		return nil, fmt.Errorf(`failed to find report %v`, err)
	}

	mappingReport := entity.ReportBuilder{
		Name:        model.Name,
		Query:       model.Query,
		RouteName:   model.RouteName,
		Description: model.Description,
		Columns:     model.Columns,
		IsActive:    model.IsActive,
	}

	return &mappingReport, nil
}

func (r *ReportBuilderRepositoryImpl) DeleteReport(reportId int) error {

	var model model.ReportBuilderModel

	if err := r.db.Where(`report_id = ?`, reportId).
		Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *ReportBuilderRepositoryImpl) UpdateReport(reportId int, item entity.ReportBuilder) error {
	var reportBuilder model.ReportBuilderModel

	payloadUpdate := model.ReportBuilderModel{
		Name:        item.Name,
		Query:       item.Query,
		RouteName:   item.RouteName,
		Description: item.Description,
		Columns:     item.Columns,
		IsActive:    item.IsActive,
	}

	if err := r.db.Model(&reportBuilder).
		Where(`report_id = ?`, reportId).
		Updates(payloadUpdate).Error; err != nil {
		return err
	}

	return nil
}
