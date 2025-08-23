package services

import (
	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/repository"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

func GetAllComponents(input models.GetAllComponentsInput) ([]models.Component, error) {
	page := input.Page
	utils.Log(constants.SERVICE_GET_ALL_COMPONENTS_START, nil, page)

	components, err := repository.GetAllComponents(input)
	if err != nil {
		utils.Log(constants.SERVICE_GET_ALL_COMPONENTS_ERROR, err)
		return nil, err
	}

	utils.Log(constants.SERVICE_GET_ALL_COMPONENTS_SUCCESS, nil, page)
	return components, nil
}

func GetComponentsByCategory(input models.GetComponentsByCategoryInput) ([]models.Component, error) {
	category, page := input.Category, input.Page
	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_CATEGORY_START, nil, category, page)

	components, err := repository.GetComponentsByCategory(input)
	if err != nil {
		utils.Log(constants.SERVICE_GET_COMPONENTS_BY_CATEGORY_ERROR, err, category, page)
		return nil, err
	}

	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_CATEGORY_SUCCESS, nil, category, page)
	return components, nil
}

func GetComponentsByBrand(input models.GetComponentsByBrandInput) ([]models.Component, error) {
	category, brand, page := input.Category, input.Brand, input.Page
	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_BRAND_START, nil, category, brand, page)

	components, err := repository.GetComponentsByBrand(input)
	if err != nil {
		utils.Log(constants.SERVICE_GET_COMPONENTS_BY_BRAND_ERROR, err, category, brand, page)
		return nil, err
	}

	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_BRAND_SUCCESS, nil, category, brand, page)
	return components, nil
}

func GetComponentById(input models.GetComponentByIdInput) (models.Component, error) {
	id, page := input.ID, input.Page
	utils.Log(constants.SERVICE_GET_COMPONENT_BY_ID_START, nil, id, page)

	component, err := repository.GetComponentById(input)
	if err != nil {
		utils.Log(constants.SERVICE_GET_COMPONENT_BY_ID_ERROR, err, id, page)
		return models.Component{}, err
	}

	utils.Log(constants.SERVICE_GET_COMPONENT_BY_ID_SUCCESS, nil, id, page)
	return component, nil
}
