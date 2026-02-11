package helpers

import (
	"errors"
	"net/http"

	domainErrors "github.com/rafmasloman/mis-modernisasi-backend/internal/delivery/errors"
)

func MapErrorToHttpStatus(err error) (statusCode int, message string) {
	switch {
	case errors.Is(err, domainErrors.ErrFilterNotFound):
		return http.StatusNotFound, "Filter not found"
	// Tambahkan error domain lainnya di sini
	case errors.Is(err, domainErrors.ErrReportNotFound):
		return http.StatusNotFound, "Report not found"
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}
