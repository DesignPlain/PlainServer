package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontroller struct {
	/*
	   Configuration of Policy Controller
	   Structure is documented below.
	*/
	PolicyControllerHubConfig Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig `json:"policyControllerHubConfig,omitempty" yaml:"policyControllerHubConfig,omitempty"`

	// Configures the version of Policy Controller
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
