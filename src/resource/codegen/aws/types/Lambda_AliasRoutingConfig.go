package types

type Lambda_AliasRoutingConfig struct {
	// A map that defines the proportion of events that should be sent to different versions of a lambda function.
	AdditionalVersionWeights map[string]float64 `json:"additionalVersionWeights,omitempty" yaml:"additionalVersionWeights,omitempty"`
}
