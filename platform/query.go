package platform

type Query struct {
	QueryFilter QueryFilter `json:"QueryFilter"`
	QuerySort   QuerySort   `json:"QuerySort"`
}

type QueryFilter struct {
	GroupingExpression GroupingExpression `json:"expression"`
}

type QuerySort struct {
	SortField []SortField `json:"sortField"`
}

type GroupingExpression struct {
	Operator         string             `json:"operator"`
	NestedExpression []SimpleExpression `json:"nestedExpression"`
}

type SimpleExpression struct {
	Property string   `json:"property"`
	Operator string   `json:"operator"`
	Argument []string `json:"argument"`
}

type SortField struct {
	FieldName string `json:"fieldName"`
	SortOrder string `json:"sortOrder"`
}
