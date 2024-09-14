package types

type Networkservices_HttpRouteRuleMatchHeaderRangeMatch struct {
	// Start of the range (inclusive).
	Start int `json:"start,omitempty" yaml:"start,omitempty"`

	// End of the range (exclusive).
	End int `json:"end,omitempty" yaml:"end,omitempty"`
}
