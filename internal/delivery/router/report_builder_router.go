package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/http"
)

func RegisterReportBuilderRoute(r *gin.RouterGroup, controller *http.ReportBuilderController) {
	reportRouter := r.Group("/report-builder")

	reportRouter.GET("/:routeName", controller.GetReportByRouterName)
	reportRouter.GET("/", controller.GetAllReportBuilder)
	reportRouter.POST("/", controller.CreateReportBuilder)
	reportRouter.PUT("/:id", controller.UpdateReportBuilder)
	reportRouter.DELETE("/:id", controller.DeleteReportBuilder)

}
