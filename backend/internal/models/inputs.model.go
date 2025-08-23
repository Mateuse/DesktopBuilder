package models

type GenerateSelectQueryInput struct {
	Table       string
	Columns     []string
	WhereClause string
	Page        string
}

type GetComponentsByBrandInput struct {
	Category string
	Brand    string
	Page     string
}

type ComponentQueryParams struct {
	Category string
	Brand    string
	ID       string
}

type GetComponentsByCategoryInput struct {
	Category string
	Page     string
}

type GetAllComponentsInput struct {
	Page string
}

type GetComponentByIdInput struct {
	ID   string
	Page string
}
