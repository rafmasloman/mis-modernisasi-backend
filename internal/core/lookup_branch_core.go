package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/router"
)

func RegisterLookupBranchCore(api *gin.RouterGroup, repos *RepositoryContainer) {

	lookupBranchUsecase := usecase.NewLookupBranchUsecase(repos.LookupBranch)
	lookupBranchController := http.NewLookupBranchController(lookupBranchUsecase)

	router.RegisterLookupBranchRoute(api, lookupBranchController)
}
