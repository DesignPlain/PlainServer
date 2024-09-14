package waf

import types "libds/aws/types"

type GeoMatchSet struct {
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []types.Waf_GeoMatchSetGeoMatchConstraint `json:"geoMatchConstraints,omitempty" yaml:"geoMatchConstraints,omitempty"`

	// The name or description of the GeoMatchSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
