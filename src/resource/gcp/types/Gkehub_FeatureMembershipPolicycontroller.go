package types

type Gkehub_FeatureMembershipPolicycontroller struct {
	// Version of Policy Controller to install. Defaults to the latest version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Policy Controller configuration for the cluster. Structure is documented below.
	PolicyControllerHubConfig Gkehub_FeatureMembershipPolicycontrollerPolicyControllerHubConfig `json:"policyControllerHubConfig,omitempty" yaml:"policyControllerHubConfig,omitempty"`
}
