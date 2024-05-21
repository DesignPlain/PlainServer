package types

type Networkservices_HttpRouteRuleMatchHeaderRangeMatch struct {
	// End of the range (exclusive).
	End int `json:"end,omitempty" yaml:"end,omitempty"`

	// Start of the range (inclusive).
	Start int `json:"start,omitempty" yaml:"start,omitempty"`
}
