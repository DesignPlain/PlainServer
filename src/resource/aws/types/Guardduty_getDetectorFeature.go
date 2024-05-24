package types

type Guardduty_getDetectorFeature struct {
	// The name of the detector feature.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Current status of the detector.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Additional feature configuration.
	AdditionalConfigurations []Guardduty_getDetectorFeatureAdditionalConfiguration `json:"additionalConfigurations,omitempty" yaml:"additionalConfigurations,omitempty"`
}
