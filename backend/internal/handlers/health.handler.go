package handlers

import (
	"net/http"

	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, constants.METHOD_NOT_ALLOWED_MESSAGE, nil)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, constants.HEALTH_MESSAGE, nil)
}
