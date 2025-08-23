package models

type GenerateSelectQueryInput struct {
	Table       string
	Columns     []string
	WhereClause string
}

type GetComponentsByBrandInput struct {
	Category string
	Brand    string
}

type ComponentQueryParams struct {
	Category string
	Brand    string
	ID       string
}
