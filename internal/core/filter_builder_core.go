package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/router"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/repositories"
	"gorm.io/gorm"
)

func RegisterFilterBuilderCore(api *gin.RouterGroup, db *gorm.DB) {

	filterBuilderRepository := repositories.NewFilterBuilderRepository(db)
	filterBuilderUsecase := usecase.NewFilterBuilderUsecase(filterBuilderRepository)
	filterBuilderController := http.NewFilterBuilderController(filterBuilderUsecase)

	router.RegisterFilterBuilderRoute(api, filterBuilderController)
}
