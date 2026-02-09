package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/dto"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/shared/helpers"
)

type ReportBuilderController struct {
	usecase usecase.ReportBuilderUsecaseImpl
}

func NewReportBuilderController(usecase usecase.ReportBuilderUsecaseImpl) *ReportBuilderController {
	return &ReportBuilderController{usecase: usecase}
}

func (c *ReportBuilderController) CreateReportBuilder(ctx *gin.Context) {
	var payload dto.DTOCreateReportBuilder
	response := new(dto.ApiResponseDTO)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponseDTO{
			ResponseCode: http.StatusBadRequest,
			Status:       false,
			Message:      helpers.ErrFailedToCreateReport,
			Data:         nil,
			Timestamp:    time.Now(),
			Error:        err.Error(),
		})

		return
	}

	err := c.usecase.CreateReportBuilder(payload)

	if err != nil {
		log.Println("create report : %v ", err)
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToCreateReport
		response.Timestamp = time.Now()
		response.Error = err.Error()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = nil
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to create report"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}

func (c *ReportBuilderController) GetAllReportBuilder(ctx *gin.Context) {
	response := new(dto.ApiResponseDTO)

	results, err := c.usecase.GetAllReportBuilder()

	if err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToGetAllReport
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = results
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to get all report"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)

}

func (c *ReportBuilderController) GetReportByRouterName(ctx *gin.Context) {
	response := new(dto.ApiResponseDTO)
	reportName := ctx.Param("reportName")

	results, err := c.usecase.GetReportByRouterName(reportName)

	if err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToGetAllReportByRouteName
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = results
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to get report"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}

func (c *ReportBuilderController) DeleteReportBuilder(ctx *gin.Context) {
	response := new(dto.ApiResponseDTO)

	reportId := ctx.Param("reportId")

	err := c.usecase.DeleteReport(reportId)

	if err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToDeleteReport
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = nil
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to delete report"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)

}

func (c *ReportBuilderController) UpdateReportBuilder(ctx *gin.Context) {
	var payload dto.DTOCreateReportBuilder
	response := new(dto.ApiResponseDTO)

	reportId := ctx.Param("reportId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponseDTO{
			ResponseCode: http.StatusBadRequest,
			Status:       false,
			Message:      "Failed to create report",
			Data:         nil,
			Timestamp:    time.Now(),
		})

		return
	}

	err := c.usecase.UpdateReport(reportId, payload)

	if err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = "Failed to create report"
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = nil
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to create report"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}
