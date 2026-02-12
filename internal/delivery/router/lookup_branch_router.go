package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
)

func RegisterLookupBranchRoute(r *gin.RouterGroup, controller *http.LookupBranchController) {
	reportRouter := r.Group("/lookup-branch")

	reportRouter.GET("/konven", controller.GetKonven)
	reportRouter.GET("/wilayah", controller.GetWilayah)

}
