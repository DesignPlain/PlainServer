package types

type Gkehub_FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	// Configures the manner in which the template library is installed on the cluster. Must be one of `ALL`, `NOT_INSTALLED` or `INSTALLATION_UNSPECIFIED`. Defaults to `ALL`.
	Installation string `json:"installation,omitempty" yaml:"installation,omitempty"`
}
