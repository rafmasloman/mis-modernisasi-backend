package repositories

import (
	"fmt"

	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/entity"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/model"
	"gorm.io/gorm"
)

type ReportBuilderRepository struct {
	db *gorm.DB
}

type ReportBuilderRepositoryImpl interface {
	CreateReportBuilder(item entity.ReportBuilder) error
	GetAllReportBuilder() (*[]entity.ReportBuilder, error)
	GetReportByRouterName(routerName string) (*entity.ReportBuilder, error)
	DeleteReport(reportId int) error
	UpdateReport(reportId int, item entity.ReportBuilder) error
}

func NewReportBuilderRepositories(db *gorm.DB) *ReportBuilderRepository {
	return &ReportBuilderRepository{db: db}
}

func (r *ReportBuilderRepository) CreateReportBuilder(item entity.ReportBuilder) error {

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

func (r *ReportBuilderRepository) GetAllReportBuilder() (*[]entity.ReportBuilder, error) {
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

func (r *ReportBuilderRepository) GetReportByRouterName(routerName string) (*entity.ReportBuilder, error) {
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

func (r *ReportBuilderRepository) DeleteReport(reportId int) error {

	var model model.ReportBuilderModel

	if err := r.db.Where(`report_id = ?`, reportId).
		Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *ReportBuilderRepository) UpdateReport(reportId int, item entity.ReportBuilder) error {
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
