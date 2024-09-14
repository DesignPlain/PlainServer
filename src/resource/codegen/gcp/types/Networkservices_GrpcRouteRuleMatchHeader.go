package types

type Networkservices_GrpcRouteRuleMatchHeader struct {
	/*
	   The type of match.
	   Default value is `EXACT`.
	   Possible values are: `TYPE_UNSPECIFIED`, `EXACT`, `REGULAR_EXPRESSION`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Required. The value of the header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Required. The key of the header.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
