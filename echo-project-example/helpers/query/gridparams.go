package query

// GridQueryBuilder is
type GridQueryBuilder interface {
	FilterClause(f GridFilter) string
	FullQuery(g GridParams, preQuery string) string
	FilterQuery(g GridParams, preQuery string) string
	SortQuery(g GridParams) string
	SortPagingQuery(g GridParams) string
}

// GridParams is
type GridParams struct {
	Page      int         `json:"page"`
	Skip      int         `json:"skip"`
	Take      int         `json:"take"`
	PageSize  int         `json:"pageSize"`
	Sort      []GridSort  `json:"sort,omitempty"`
	Group     interface{} `json:"group,omitempty"`
	HasSort   bool        `json:"hasSort,omitempty"`
	HasFilter bool        `json:"hasFilter,omitempty"`
	Filter    struct {
		Logic   string       `json:"logic,omitempty"`
		Filters []GridFilter `json:"filters,omitempty"`
	} `json:"filter,omitempty"`
}

// GridSort is
type GridSort struct {
	Field string `json:"field"`
	Dir   string `json:"dir"`
}

// GridFilter is
type GridFilter struct {
	Field    string `json:"field,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`

	// Filter can be nested
	HasSubFilter bool         `json:"hasSubFilter,omitempty"`
	Logic        string       `json:"logic,omitempty"`
	Filters      []GridFilter `json:"filters,omitempty"`
}

// ComparisonOperator is
type ComparisonOperator struct {
	Operator       string
	WildcardBefore bool
	WildcardAfter  bool
	Unary          bool
}

var operatorMap = map[string]ComparisonOperator{
	"eq": ComparisonOperator{
		Operator: "=",
		Unary:    false,
	},
	"neq": ComparisonOperator{
		Operator: "!=",
		Unary:    false,
	},
	"contains": ComparisonOperator{
		Operator:       "LIKE",
		WildcardBefore: true,
		WildcardAfter:  true,
		Unary:          false,
	},
	"doesnotcontain": ComparisonOperator{
		Operator:       "NOT LIKE",
		WildcardBefore: true,
		WildcardAfter:  true,
		Unary:          false,
	},
	"startswith": ComparisonOperator{
		Operator:      "LIKE",
		WildcardAfter: true,
		Unary:         false,
	},
	"endswith": ComparisonOperator{
		Operator:       "LIKE",
		Unary:          false,
		WildcardBefore: true,
	},
	"isnull": ComparisonOperator{
		Operator: "IS NULL",
		Unary:    true,
	},
	"isnotnull": ComparisonOperator{
		Operator: "IS NOT NULL",
		Unary:    true,
	},
	"isempty": ComparisonOperator{
		Operator: "IS VALUED",
		Unary:    true,
	},
	"isnotempty": ComparisonOperator{
		Operator: "IS NOT VALUED",
		Unary:    true,
	},
}
