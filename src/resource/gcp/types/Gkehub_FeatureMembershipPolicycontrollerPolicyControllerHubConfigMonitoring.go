package types

type Gkehub_FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring struct {
	// Specifies the list of backends Policy Controller will export to. Must be one of `CLOUD_MONITORING` or `PROMETHEUS`. Defaults to [`CLOUD_MONITORING`, `PROMETHEUS`]. Specifying an empty value `[]` disables metrics export.
	Backends []string `json:"backends,omitempty" yaml:"backends,omitempty"`
}
