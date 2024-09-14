package types

type Cloudtrail_EventDataStoreAdvancedEventSelector struct {
	// Specifies the selector statements in an advanced event selector. Fields documented below.
	FieldSelectors []Cloudtrail_EventDataStoreAdvancedEventSelectorFieldSelector `json:"fieldSelectors,omitempty" yaml:"fieldSelectors,omitempty"`

	// Specifies the name of the advanced event selector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
