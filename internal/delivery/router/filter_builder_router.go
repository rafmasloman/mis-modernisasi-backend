package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
)

func RegisterFilterBuilderRoute(r *gin.RouterGroup, controller *http.FilterBuilderController) {
	filterRouter := r.Group("/filter-builder")

	filterRouter.POST("/", controller.CreateFilterBuilder)
	filterRouter.GET("/reportId", controller.GetFilterBuilderByReportId)
	filterRouter.PUT("/:id", controller.UpdateFilterBuilder)
	filterRouter.DELETE("/:id", controller.DeleteFilterBuilder)

}
