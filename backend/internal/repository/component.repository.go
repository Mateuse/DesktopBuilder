package repository

import (
	"fmt"

	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

func GetAllComponents(input models.GetAllComponentsInput) ([]models.Component, error) {
	page := input.Page
	utils.Log(constants.REPOSITORY_GET_ALL_COMPONENTS_START, nil, page)

	queryInput := models.GenerateSelectQueryInput{
		Table:       constants.COMPONENTS_TABLE,
		Columns:     constants.COMPONENTS_SELECT_COLUMNS,
		WhereClause: "",
	}

	query, err := utils.GenerateSelectQuery(queryInput)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_ALL_COMPONENTS_QUERY_ERROR, err)
		return nil, err
	}

	rows, err := utils.GetDB().Query(query)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_ALL_COMPONENTS_DB_ERROR, err)
		return nil, err
	}
	defer rows.Close()

	components := []models.Component{}
	for rows.Next() {
		var component models.Component
		err := rows.Scan(&component.ID, &component.Category, &component.Brand, &component.Model, &component.SKU, &component.UPC, &component.Specs, &component.CreatedAt)
		if err != nil {
			utils.Log(constants.REPOSITORY_GET_ALL_COMPONENTS_SCAN_ERROR, err)
			return nil, err
		}
		components = append(components, component)
	}

	utils.Log(constants.REPOSITORY_GET_ALL_COMPONENTS_SUCCESS, nil, len(components))
	return components, nil
}

func GetComponentsByCategory(input models.GetComponentsByCategoryInput) ([]models.Component, error) {
	category, page := input.Category, input.Page
	utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_CATEGORY_START, nil, category, page)

	whereClause := fmt.Sprintf("category = '%s'", category)
	queryInput := models.GenerateSelectQueryInput{
		Table:       constants.COMPONENTS_TABLE,
		Columns:     constants.COMPONENTS_SELECT_COLUMNS,
		WhereClause: whereClause,
	}

	query, err := utils.GenerateSelectQuery(queryInput)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_CATEGORY_QUERY_ERROR, err, category, page)
		return nil, err
	}

	rows, err := utils.GetDB().Query(query)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_CATEGORY_DB_ERROR, err, category, page)
		return nil, err
	}
	defer rows.Close()

	components := []models.Component{}
	for rows.Next() {
		var component models.Component
		err := rows.Scan(&component.ID, &component.Category, &component.Brand, &component.Model, &component.SKU, &component.UPC, &component.Specs, &component.CreatedAt)
		if err != nil {
			utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_CATEGORY_SCAN_ERROR, err, category, page)
			return nil, err
		}
		components = append(components, component)
	}

	utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_CATEGORY_SUCCESS, nil, len(components), category, page)
	return components, nil
}

func GetComponentsByBrand(input models.GetComponentsByBrandInput) ([]models.Component, error) {
	category, brand, page := input.Category, input.Brand, input.Page
	utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_BRAND_START, nil, category, brand, page)

	whereClause := fmt.Sprintf("category = '%s' AND brand = '%s'", category, brand)
	queryInput := models.GenerateSelectQueryInput{
		Table:       constants.COMPONENTS_TABLE,
		Columns:     constants.COMPONENTS_SELECT_COLUMNS,
		WhereClause: whereClause,
	}

	query, err := utils.GenerateSelectQuery(queryInput)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_BRAND_QUERY_ERROR, err, category, brand, page)
		return nil, err
	}

	rows, err := utils.GetDB().Query(query)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_BRAND_DB_ERROR, err, category, brand, page)
		return nil, err
	}
	defer rows.Close()

	components := []models.Component{}
	for rows.Next() {
		var component models.Component
		err := rows.Scan(&component.ID, &component.Category, &component.Brand, &component.Model, &component.SKU, &component.UPC, &component.Specs, &component.CreatedAt)
		if err != nil {
			utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_BRAND_SCAN_ERROR, err, category, brand, page)
			return nil, err
		}
		components = append(components, component)
	}

	utils.Log(constants.REPOSITORY_GET_COMPONENTS_BY_BRAND_SUCCESS, nil, len(components), category, brand, page)
	return components, nil
}

func GetComponentById(input models.GetComponentByIdInput) (models.Component, error) {
	id, page := input.ID, input.Page
	utils.Log(constants.REPOSITORY_GET_COMPONENT_BY_ID_START, nil, id, page)

	whereClause := fmt.Sprintf("id = '%s'", id)
	queryInput := models.GenerateSelectQueryInput{
		Table:       constants.COMPONENTS_TABLE,
		Columns:     constants.COMPONENTS_SELECT_COLUMNS,
		WhereClause: whereClause,
	}

	query, err := utils.GenerateSelectQuery(queryInput)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_COMPONENT_BY_ID_QUERY_ERROR, err, id, page)
		return models.Component{}, err
	}

	row := utils.GetDB().QueryRow(query)

	var component models.Component
	err = row.Scan(&component.ID, &component.Category, &component.Brand, &component.Model, &component.SKU, &component.UPC, &component.Specs, &component.CreatedAt)
	if err != nil {
		utils.Log(constants.REPOSITORY_GET_COMPONENT_BY_ID_SCAN_ERROR, err, id, page)
		return models.Component{}, err
	}

	utils.Log(constants.REPOSITORY_GET_COMPONENT_BY_ID_SUCCESS, nil, id, page)
	return component, nil
}
