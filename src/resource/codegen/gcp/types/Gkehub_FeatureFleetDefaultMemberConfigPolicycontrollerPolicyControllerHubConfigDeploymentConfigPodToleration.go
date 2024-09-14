package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigPodToleration struct {
	// Matches a taint effect.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// Matches a taint key (not necessarily unique).
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Matches a taint operator.
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// Matches a taint value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
