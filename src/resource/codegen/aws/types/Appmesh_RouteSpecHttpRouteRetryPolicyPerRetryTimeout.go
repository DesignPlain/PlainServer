package types

type Appmesh_RouteSpecHttpRouteRetryPolicyPerRetryTimeout struct {
	// Retry unit. Valid values: `ms`, `s`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Retry value.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
