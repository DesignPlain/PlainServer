package types

type Cloudtrail_TrailAdvancedEventSelector struct {
	// Name of the trail.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the selector statements in an advanced event selector. Fields documented below.
	FieldSelectors []Cloudtrail_TrailAdvancedEventSelectorFieldSelector `json:"fieldSelectors,omitempty" yaml:"fieldSelectors,omitempty"`
}
