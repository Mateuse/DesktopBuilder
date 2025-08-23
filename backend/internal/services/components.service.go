package services

import (
	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/repository"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

func GetAllComponents() ([]models.Component, error) {
	utils.Log(constants.SERVICE_GET_ALL_COMPONENTS_START, nil)

	components, err := repository.GetAllComponents()
	if err != nil {
		utils.Log(constants.SERVICE_GET_ALL_COMPONENTS_ERROR, err)
		return nil, err
	}

	utils.Log(constants.SERVICE_GET_ALL_COMPONENTS_SUCCESS, nil)
	return components, nil
}

func GetComponentsByCategory(category string) ([]models.Component, error) {
	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_CATEGORY_START, nil, category)

	components, err := repository.GetComponentsByCategory(category)
	if err != nil {
		utils.Log(constants.SERVICE_GET_COMPONENTS_BY_CATEGORY_ERROR, err, category)
		return nil, err
	}

	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_CATEGORY_SUCCESS, nil, category)
	return components, nil
}

func GetComponentsByBrand(input models.GetComponentsByBrandInput) ([]models.Component, error) {
	category, brand := input.Category, input.Brand
	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_BRAND_START, nil, category, brand)

	components, err := repository.GetComponentsByBrand(category, brand)
	if err != nil {
		utils.Log(constants.SERVICE_GET_COMPONENTS_BY_BRAND_ERROR, err, category, brand)
		return nil, err
	}

	utils.Log(constants.SERVICE_GET_COMPONENTS_BY_BRAND_SUCCESS, nil, category, brand)
	return components, nil
}

func GetComponentById(id string) (models.Component, error) {
	utils.Log(constants.SERVICE_GET_COMPONENT_BY_ID_START, nil, id)

	component, err := repository.GetComponentById(id)
	if err != nil {
		utils.Log(constants.SERVICE_GET_COMPONENT_BY_ID_ERROR, err, id)
		return models.Component{}, err
	}

	utils.Log(constants.SERVICE_GET_COMPONENT_BY_ID_SUCCESS, nil, id)
	return component, nil
}
