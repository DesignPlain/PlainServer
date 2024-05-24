package types

type Wafregional_GeoMatchSetGeoMatchConstraint struct {
	// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The country that you want AWS WAF to search for.
	   This is the two-letter country code, e.g., `US`, `CA`, `RU`, `CN`, etc.
	   See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchConstraint.html) for all supported values.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
