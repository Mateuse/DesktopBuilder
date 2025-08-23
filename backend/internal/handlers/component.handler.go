package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/services"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

func GetComponentsHandler(w http.ResponseWriter, r *http.Request) {
	utils.Log(constants.HANDLER_GET_COMPONENTS_START, nil)

	if r.Method != http.MethodGet {
		utils.Log(constants.HANDLER_METHOD_NOT_ALLOWED, fmt.Errorf("method %s not allowed", r.Method))
		utils.WriteError(w, http.StatusMethodNotAllowed, constants.METHOD_NOT_ALLOWED_MESSAGE, nil)
		return
	}

	params := parseComponentQueryParams(r)

	switch {
	case params.ID != "":
		handleGetComponentByID(w, params.ID)
	case params.Category != "" && params.Brand != "":
		handleGetComponentsByBrand(w, params.Category, params.Brand)
	case params.Category != "":
		handleGetComponentsByCategory(w, params.Category)
	default:
		handleGetAllComponents(w)
	}
}

func parseComponentQueryParams(r *http.Request) models.ComponentQueryParams {
	brand := r.PathValue("brand")
	if brand != "" {
		brand = strings.ToLower(brand)
	}

	return models.ComponentQueryParams{
		Category: r.PathValue("category"),
		Brand:    brand,
		ID:       r.PathValue("id"),
	}
}

func handleGetComponentByID(w http.ResponseWriter, id string) {
	utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_START, nil, id)

	component, err := services.GetComponentById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_NOT_FOUND, nil, id)
			utils.WriteError(w, http.StatusNotFound, constants.COMPONENT_NOT_FOUND_MESSAGE, nil)
			return
		}
		utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_ERROR, err, id)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_SUCCESS, nil, id)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, component)
}

func handleGetComponentsByBrand(w http.ResponseWriter, category, brand string) {
	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_BRAND_START, nil, category, brand)

	input := models.GetComponentsByBrandInput{
		Category: category,
		Brand:    brand,
	}
	components, err := services.GetComponentsByBrand(input)
	if err != nil {
		utils.Log(constants.HANDLER_GET_COMPONENTS_BY_BRAND_ERROR, err, category, brand)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_BRAND_SUCCESS, nil, category, brand)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, components)
}

func handleGetComponentsByCategory(w http.ResponseWriter, category string) {
	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_CATEGORY_START, nil, category)

	components, err := services.GetComponentsByCategory(category)
	if err != nil {
		utils.Log(constants.HANDLER_GET_COMPONENTS_BY_CATEGORY_ERROR, err, category)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_CATEGORY_SUCCESS, nil, category)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, components)
}

func handleGetAllComponents(w http.ResponseWriter) {
	utils.Log(constants.HANDLER_GET_ALL_COMPONENTS_START, nil)

	components, err := services.GetAllComponents()
	if err != nil {
		utils.Log(constants.HANDLER_GET_ALL_COMPONENTS_ERROR, err)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_ALL_COMPONENTS_SUCCESS, nil)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, components)
}
