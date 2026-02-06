package dto

import "time"

type ApiResponseDTO struct {
	ResponseCode int         `json:"response_code"`
	Status       bool        `json:"status" from:"status"`
	Message      string      `json:"message" from:"message"`
	Data         any         `json:"data"`
	Error        interface{} `json:"error"`
	Timestamp    time.Time   `json:"timestamp"`
}
