package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/router"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/repositories"
	"gorm.io/gorm"
)

func RegisterReportBuilderCore(api *gin.RouterGroup, db *gorm.DB) {

	reportBuilderRepository := repositories.NewReportBuilderRepositories(db)
	reportBuilderUsecase := usecase.NewReportBuilderUsecase(reportBuilderRepository)
	reportBuilderController := http.NewReportBuilderController(reportBuilderUsecase)

	router.RegisterReportBuilderRoute(api, reportBuilderController)

}
