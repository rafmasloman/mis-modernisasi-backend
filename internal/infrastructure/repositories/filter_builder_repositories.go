package repositories

import (
	"fmt"

	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/entity"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/model"
	"gorm.io/gorm"
)

type FilterBuilderRepository struct {
	db *gorm.DB
}

type FilterBuilderRepositoryImpl interface {
	CreateFilterBuilder(item entity.FilterBuilder) error
	GetFilterBuilderByReportId(reportId int) (*[]entity.FilterBuilder, error)
	DeleteFilterBuilder(id int) error
	UpdateFilterBuilder(id int, item entity.FilterBuilder) error
}

func NewFilterBuilderRepository(db *gorm.DB) *FilterBuilderRepository {
	return &FilterBuilderRepository{db: db}
}

func (r *FilterBuilderRepository) CreateFilterBuilder(item entity.FilterBuilder) error {

	reportData := model.FilterBuilderModel{
		Name:     item.Name,
		Title:    item.Title,
		Query:    item.Query,
		Type:     item.Type,
		Required: item.Required,
		ReportId: item.ReportId,
		OrderNum: item.OrderNum,
	}

	if err := r.db.Create(&reportData).Error; err != nil {
		return err
	}

	return nil
}

func (r *FilterBuilderRepository) GetFilterBuilderByReportId(reportId int) (*[]entity.FilterBuilder, error) {
	var model []model.FilterBuilderModel

	if err := r.db.Table(`report_builder`).Where(`route_name`).First(&model).Error; err != nil {
		return nil, fmt.Errorf(`failed to find filter %v`, err)
	}

	entities := make([]entity.FilterBuilder, 0, len(model))

	for _, item := range model {
		entities = append(entities, entity.FilterBuilder{
			Name:     item.Name,
			Query:    item.Query,
			Title:    item.Title,
			ReportId: item.ReportId,
			Required: item.Required,
			Type:     item.Type,
			OrderNum: item.OrderNum,
		})

	}

	return &entities, nil
}

func (r *FilterBuilderRepository) DeleteFilterBuilder(id int) error {
	var model model.FilterBuilderModel

	if err := r.db.Where(`id = ?`, id).
		Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *FilterBuilderRepository) UpdateFilterBuilder(id int, item entity.FilterBuilder) error {
	var filterBuilder model.FilterBuilderModel

	payloadUpdate := model.FilterBuilderModel{
		Name:     item.Name,
		Title:    item.Title,
		Query:    item.Query,
		ReportId: item.ReportId,
		Type:     item.Type,
		Required: item.Required,
		OrderNum: item.OrderNum,
	}

	if err := r.db.Model(&filterBuilder).
		Where(`id = ?`, id).
		Updates(payloadUpdate).Error; err != nil {
		return err
	}

	return nil
}
