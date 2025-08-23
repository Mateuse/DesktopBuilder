package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

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
	page := utils.GetPageNumberFromQueryString(r.URL.Query())

	switch {
	case params.ID != "":
		input := models.GetComponentByIdInput{
			ID:   params.ID,
			Page: page,
		}
		handleGetComponentByID(w, input)
	case params.Category != "" && params.Brand != "":
		input := models.GetComponentsByBrandInput{
			Category: params.Category,
			Brand:    params.Brand,
			Page:     page,
		}
		handleGetComponentsByBrand(w, input)
	case params.Category != "":
		input := models.GetComponentsByCategoryInput{
			Category: params.Category,
			Page:     page,
		}
		handleGetComponentsByCategory(w, input)
	default:
		input := models.GetAllComponentsInput{
			Page: page,
		}
		handleGetAllComponents(w, input)
	}
}

func parseComponentQueryParams(r *http.Request) models.ComponentQueryParams {
	brand := r.PathValue("brand")

	return models.ComponentQueryParams{
		Category: r.PathValue("category"),
		Brand:    brand,
		ID:       r.PathValue("id"),
	}
}

func handleGetComponentByID(w http.ResponseWriter, input models.GetComponentByIdInput) {
	utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_START, nil, input.ID)

	component, err := services.GetComponentById(input)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_NOT_FOUND, nil, input.ID)
			utils.WriteError(w, http.StatusNotFound, constants.COMPONENT_NOT_FOUND_MESSAGE, nil)
			return
		}
		utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_ERROR, err, input.ID)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_COMPONENT_BY_ID_SUCCESS, nil, input.ID)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, component)
}

func handleGetComponentsByBrand(w http.ResponseWriter, input models.GetComponentsByBrandInput) {
	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_BRAND_START, nil, input.Category, input.Brand)

	components, err := services.GetComponentsByBrand(input)
	if err != nil {
		utils.Log(constants.HANDLER_GET_COMPONENTS_BY_BRAND_ERROR, err, input.Category, input.Brand)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_BRAND_SUCCESS, nil, input.Category, input.Brand)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, components)
}

func handleGetComponentsByCategory(w http.ResponseWriter, input models.GetComponentsByCategoryInput) {
	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_CATEGORY_START, nil, input.Category)

	components, err := services.GetComponentsByCategory(input)
	if err != nil {
		utils.Log(constants.HANDLER_GET_COMPONENTS_BY_CATEGORY_ERROR, err, input.Category)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_COMPONENTS_BY_CATEGORY_SUCCESS, nil, input.Category)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, components)
}

func handleGetAllComponents(w http.ResponseWriter, input models.GetAllComponentsInput) {
	utils.Log(constants.HANDLER_GET_ALL_COMPONENTS_START, nil)

	components, err := services.GetAllComponents(input)
	if err != nil {
		utils.Log(constants.HANDLER_GET_ALL_COMPONENTS_ERROR, err)
		utils.WriteError(w, http.StatusInternalServerError, constants.INTERNAL_SERVER_ERROR_MESSAGE, err)
		return
	}

	utils.Log(constants.HANDLER_GET_ALL_COMPONENTS_SUCCESS, nil)
	utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, components)
}
