package types

type Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig struct {
	// Logs all denies and dry run failures.
	LogDeniesEnabled bool `json:"logDeniesEnabled,omitempty" yaml:"logDeniesEnabled,omitempty"`

	/*
	   Specifies the desired policy content on the cluster.
	   Structure is documented below.
	*/
	PolicyContent Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent `json:"policyContent,omitempty" yaml:"policyContent,omitempty"`

	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	ReferentialRulesEnabled bool `json:"referentialRulesEnabled,omitempty" yaml:"referentialRulesEnabled,omitempty"`

	/*
	   Map of deployment configs to deployments ("admission", "audit", "mutation").
	   Structure is documented below.
	*/
	DeploymentConfigs []Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfig `json:"deploymentConfigs,omitempty" yaml:"deploymentConfigs,omitempty"`

	// The maximum number of audit violations to be stored in a constraint. If not set, the internal default of 20 will be used.
	ConstraintViolationLimit int `json:"constraintViolationLimit,omitempty" yaml:"constraintViolationLimit,omitempty"`

	// The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty" yaml:"exemptableNamespaces,omitempty"`

	/*
	   Configures the mode of the Policy Controller installation
	   Possible values are: `INSTALL_SPEC_UNSPECIFIED`, `INSTALL_SPEC_NOT_INSTALLED`, `INSTALL_SPEC_ENABLED`, `INSTALL_SPEC_SUSPENDED`, `INSTALL_SPEC_DETACHED`.
	*/
	InstallSpec string `json:"installSpec,omitempty" yaml:"installSpec,omitempty"`

	/*
	   Monitoring specifies the configuration of monitoring Policy Controller.
	   Structure is documented below.
	*/
	Monitoring Gkehub_FeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	// Enables the ability to mutate resources using Policy Controller.
	MutationEnabled bool `json:"mutationEnabled,omitempty" yaml:"mutationEnabled,omitempty"`

	// Interval for Policy Controller Audit scans (in seconds). When set to 0, this disables audit functionality altogether.
	AuditIntervalSeconds int `json:"auditIntervalSeconds,omitempty" yaml:"auditIntervalSeconds,omitempty"`
}
