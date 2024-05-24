package types

type Guardduty_getDetectorFeatureAdditionalConfiguration struct {
	// The name of the detector feature.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Current status of the detector.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
