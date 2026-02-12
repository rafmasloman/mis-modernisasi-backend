package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/usecase"
	responseHelpers "github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/helpers"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/shared/helpers"
)

type LookupBranchController struct {
	usecase usecase.LookupBranchUsecase
}

func NewLookupBranchController(usecase usecase.LookupBranchUsecase) *LookupBranchController {
	return &LookupBranchController{usecase: usecase}
}

func (c *LookupBranchController) GetKonven(ctx *gin.Context) {
	results, err := c.usecase.GetKonven()

	if err != nil {
		responseHelpers.NewResponse(ctx).
			Error(http.StatusBadRequest, helpers.ErrFailedToGetRegion).Send()

		return
	}

	responseHelpers.NewResponse(ctx).
		WithData(results).
		Success("Success to get flag konven").
		Send()
}

func (c *LookupBranchController) GetWilayah(ctx *gin.Context) {
	flag := ctx.Query("flag")

	results, err := c.usecase.GetWilayah(flag)

	if err != nil {
		responseHelpers.NewResponse(ctx).
			Error(http.StatusBadRequest, helpers.ErrFailedToGetRegion).Send()

		return
	}

	responseHelpers.NewResponse(ctx).
		WithData(results).
		Success("Success to get flag konven").
		Send()
}
