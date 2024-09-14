package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontroller struct {
	// Configures the version of Policy Controller
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   Configuration of Policy Controller
	   Structure is documented below.
	*/
	PolicyControllerHubConfig Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig `json:"policyControllerHubConfig,omitempty" yaml:"policyControllerHubConfig,omitempty"`
}
