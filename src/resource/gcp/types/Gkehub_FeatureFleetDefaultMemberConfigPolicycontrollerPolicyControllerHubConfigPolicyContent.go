package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent struct {
	/*
	   Configures which bundles to install and their corresponding install specs.
	   Structure is documented below.
	*/
	Bundles []Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentBundle `json:"bundles,omitempty" yaml:"bundles,omitempty"`

	/*
	   Configures the installation of the Template Library.
	   Structure is documented below.
	*/
	TemplateLibrary Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary `json:"templateLibrary,omitempty" yaml:"templateLibrary,omitempty"`
}
