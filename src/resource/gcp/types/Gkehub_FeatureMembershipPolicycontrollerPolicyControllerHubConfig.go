package types

type Gkehub_FeatureMembershipPolicycontrollerPolicyControllerHubConfig struct {
	// The maximum number of audit violations to be stored in a constraint. If not set, the  default of 20 will be used.
	ConstraintViolationLimit int `json:"constraintViolationLimit,omitempty" yaml:"constraintViolationLimit,omitempty"`

	// The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty" yaml:"exemptableNamespaces,omitempty"`

	// Configures the mode of the Policy Controller installation. Must be one of `INSTALL_SPEC_NOT_INSTALLED`, `INSTALL_SPEC_ENABLED`, `INSTALL_SPEC_SUSPENDED` or `INSTALL_SPEC_DETACHED`.
	InstallSpec string `json:"installSpec,omitempty" yaml:"installSpec,omitempty"`

	// Logs all denies and dry run failures.
	LogDeniesEnabled bool `json:"logDeniesEnabled,omitempty" yaml:"logDeniesEnabled,omitempty"`

	// Specifies the desired policy content on the cluster. Structure is documented below.
	PolicyContent Gkehub_FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent `json:"policyContent,omitempty" yaml:"policyContent,omitempty"`

	// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
	AuditIntervalSeconds int `json:"auditIntervalSeconds,omitempty" yaml:"auditIntervalSeconds,omitempty"`

	// Specifies the backends Policy Controller should export metrics to. Structure is documented below.
	Monitoring Gkehub_FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	// Enables mutation in policy controller. If true, mutation CRDs, webhook, and controller deployment will be deployed to the cluster.
	MutationEnabled bool `json:"mutationEnabled,omitempty" yaml:"mutationEnabled,omitempty"`

	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	ReferentialRulesEnabled bool `json:"referentialRulesEnabled,omitempty" yaml:"referentialRulesEnabled,omitempty"`
}
