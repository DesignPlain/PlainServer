package wafregional

import types "libds/aws/types"

type GeoMatchSet struct {
	// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []types.Wafregional_GeoMatchSetGeoMatchConstraint `json:"geoMatchConstraints,omitempty" yaml:"geoMatchConstraints,omitempty"`

	// The name or description of the Geo Match Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
