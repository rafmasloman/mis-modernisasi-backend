package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/router"
	"gorm.io/gorm"
)

func RegisterReportBuilderCore(api *gin.RouterGroup, db *gorm.DB, repos *RepositoryContainer) {

	reportBuilderUsecase := usecase.NewReportBuilderUsecase(repos.ReportBuilder)
	reportBuilderController := http.NewReportBuilderController(reportBuilderUsecase)

	router.RegisterReportBuilderRoute(api, reportBuilderController)

}
