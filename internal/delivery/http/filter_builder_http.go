package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/dto"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	responseHelpers "github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/helpers"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/shared/helpers"
)

type FilterBuilderController struct {
	usecase usecase.FilterBuilderUsecase
}

func NewFilterBuilderController(usecase usecase.FilterBuilderUsecase) *FilterBuilderController {
	return &FilterBuilderController{usecase: usecase}
}

func (c *FilterBuilderController) GetAllFilterBuilder(ctx *gin.Context) {
	response := new(dto.ApiResponseDTO)

	results, err := c.usecase.GetAllFilterBuilder()

	if err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToCreateFilter
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = results
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to get all filter"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}

func (c *FilterBuilderController) CreateFilterBuilder(ctx *gin.Context) {
	var payload dto.FilterBuilderDTO

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		responseHelpers.NewResponse(ctx).
			Error(http.StatusBadRequest, helpers.ErrFailedFilterPayload).
			Send()

		return
	}

	err := c.usecase.CreateFilterBuilder(payload)

	if err != nil {
		statusCode, message := responseHelpers.MapErrorToHttpStatus(err)
		responseHelpers.NewResponse(ctx).
			Error(statusCode, message).
			Send()

		return
	}

	responseHelpers.NewResponse(ctx).
		Success("Success to create filter builder").
		Send()
}

func (c *FilterBuilderController) GetFilterBuilderByReportId(ctx *gin.Context) {
	response := new(dto.ApiResponseDTO)

	reportId := ctx.Param("reportId")

	results, err := c.usecase.GetFilterBuilderByReportId(reportId)

	if err != nil {
		log.Println("create report : %v ", err)
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToGetFilterByReportId
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.Data = results
	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to get filter builder by report id"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}

func (c *FilterBuilderController) DeleteFilterBuilder(ctx *gin.Context) {

	id := ctx.Query("id")

	response := new(dto.ApiResponseDTO)

	if err := c.usecase.DeleteFilterBuilder(id); err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToGetAllFilter
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to delete filter"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}

func (c *FilterBuilderController) UpdateFilterBuilder(ctx *gin.Context) {

	var payload dto.FilterBuilderDTO
	response := new(dto.ApiResponseDTO)

	id := ctx.Query("id")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ApiResponseDTO{
			ResponseCode: http.StatusBadRequest,
			Status:       false,
			Message:      helpers.ErrFailedToCreateReport,
			Data:         nil,
			Timestamp:    time.Now(),
		})

		return
	}

	err := c.usecase.UpdateFilterBuilder(id, payload)

	if err != nil {
		response.Data = nil
		response.ResponseCode = http.StatusBadGateway
		response.Status = false
		response.Message = helpers.ErrFailedToGetAllFilter
		response.Timestamp = time.Now()

		ctx.JSON(http.StatusBadGateway, response)

		return
	}

	response.ResponseCode = http.StatusOK
	response.Status = true
	response.Message = "Success to update filter builder"
	response.Timestamp = time.Now()

	ctx.JSON(http.StatusOK, response)
}
