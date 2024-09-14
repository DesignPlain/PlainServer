package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentBundle struct {
	// The identifier for this object. Format specified above.
	Bundle string `json:"bundle,omitempty" yaml:"bundle,omitempty"`

	// The set of namespaces to be exempted from the bundle.
	ExemptedNamespaces []string `json:"exemptedNamespaces,omitempty" yaml:"exemptedNamespaces,omitempty"`
}
