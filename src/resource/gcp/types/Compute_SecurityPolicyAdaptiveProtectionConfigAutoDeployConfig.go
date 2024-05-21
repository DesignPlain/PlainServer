package types

type Compute_SecurityPolicyAdaptiveProtectionConfigAutoDeployConfig struct {
	// Rules are only automatically deployed for alerts on potential attacks with confidence scores greater than this threshold.
	ConfidenceThreshold float64 `json:"confidenceThreshold,omitempty" yaml:"confidenceThreshold,omitempty"`

	// Google Cloud Armor stops applying the action in the automatically deployed rule to an identified attacker after this duration. The rule continues to operate against new requests.
	ExpirationSec int `json:"expirationSec,omitempty" yaml:"expirationSec,omitempty"`

	// Rules are only automatically deployed when the estimated impact to baseline traffic from the suggested mitigation is below this threshold.
	ImpactedBaselineThreshold float64 `json:"impactedBaselineThreshold,omitempty" yaml:"impactedBaselineThreshold,omitempty"`

	// Identifies new attackers only when the load to the backend service that is under attack exceeds this threshold.
	LoadThreshold float64 `json:"loadThreshold,omitempty" yaml:"loadThreshold,omitempty"`
}
