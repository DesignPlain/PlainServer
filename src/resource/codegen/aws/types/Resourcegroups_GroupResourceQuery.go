package types

type Resourcegroups_GroupResourceQuery struct {
	// The resource query as a JSON string.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`

	// The type of the resource query. Defaults to `TAG_FILTERS_1_0`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
