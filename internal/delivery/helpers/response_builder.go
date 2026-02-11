package helpers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/application/dto"
)

type ResponseBuilder struct {
	ctx        *gin.Context
	statusCode int
	message    string
	data       interface{}
}

func NewResponse(ctx *gin.Context) *ResponseBuilder {
	return &ResponseBuilder{ctx: ctx}
}

func (r *ResponseBuilder) Success(message string) *ResponseBuilder {
	r.statusCode = http.StatusOK
	r.message = message

	return r
}

func (r *ResponseBuilder) Error(statusCode int, message string) *ResponseBuilder {
	r.statusCode = statusCode
	r.message = message

	return r
}

func (r *ResponseBuilder) WithData(data interface{}) *ResponseBuilder {
	r.data = data
	return r
}

func (r *ResponseBuilder) Send() {
	r.ctx.JSON(r.statusCode, dto.ApiResponseDTO{
		ResponseCode: r.statusCode,
		Status:       r.statusCode >= 200 && r.statusCode < 300,
		Message:      r.message,
		Data:         r.data,
		Timestamp:    time.Now(),
	})
}
