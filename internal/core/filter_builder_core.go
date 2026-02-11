package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/router"
	"gorm.io/gorm"
)

func RegisterFilterBuilderCore(api *gin.RouterGroup, db *gorm.DB, repos *RepositoryContainer) {

	filterBuilderUsecase := usecase.NewFilterBuilderUsecase(repos.FilterBuilder, repos.ReportBuilder)
	filterBuilderController := http.NewFilterBuilderController(filterBuilderUsecase)

	router.RegisterFilterBuilderRoute(api, filterBuilderController)
}
