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

type Folder struct {
	PermittedRoles []RoleReference `json:"PermittedRoles"`
	Deleted        bool            `json:"deleted"`
	FullPath       string          `json:"fullPath"`
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	ParentId       string          `json:"parentId"`
	ParentName     string          `json:"parentName"`
}

type RoleReference struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func QueryWithSimpleFilter(property string, operator string, argument string) Query {
	return Query{
		QueryFilter: QueryFilter{
			GroupingExpression: GroupingExpression{
				Operator:         "AND",
				NestedExpression: []SimpleExpression{{Property: property, Operator: operator, Argument: []string{argument}}},
			},
		},
	}
}

type QueryResponse struct {
	Type       string `json:"@type"`
	QueryToken string `json:"queryToken"`
	Result     any    `json:"result"`
	// Result []@type (Array of objects of the type specified in the @type attribute)
	// manually unmarshall and type assert
	NumberOfResults int `json:"numberOfResults"`
}

type QueryFolderResponse struct {
	Type            string    `json:"@type"`
	QueryToken      string    `json:"queryToken"`
	Result          []*Folder `json:"result"`
	NumberOfResults int       `json:"numberOfResults"`
}

type QueryComponentMetadataResponse struct {
	Type            string               `json:"@type"`
	QueryToken      string               `json:"queryToken"`
	Result          []*ComponentMetadata `json:"result"`
	NumberOfResults int                  `json:"numberOfResults"`
}

type QueryDeployedPackageResponse struct {
	Type            string             `json:"@type"`
	QueryToken      string             `json:"queryToken"`
	Result          []*DeployedPackage `json:"result"`
	NumberOfResults int                `json:"numberOfResults"`
}
