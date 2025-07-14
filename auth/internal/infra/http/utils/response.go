package utils

import (
	"encoding/json"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
	"github.com/dawit_hopes/saas/auth/internal/infra/log"

	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, error appErr.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.Code)

	response := map[string]string{
		"error": error.Message,
	}

	pretty, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Logger.Error("failed to encode error response: " + err.Error())
		return
	}

	w.Write(pretty)
}

func SendSuccessResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Logger.Error("failed to encode success response: " + err.Error())
		return
	}

	w.Write(pretty)
}
