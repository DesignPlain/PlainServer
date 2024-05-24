package types

type Resourceexplorer_ViewFilters struct {
	// The string that contains the search keywords, prefixes, and operators to control the results that can be returned by a search operation. For more details, see [Search query syntax](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html).
	FilterString string `json:"filterString,omitempty" yaml:"filterString,omitempty"`
}
