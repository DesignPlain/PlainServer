package types

type Gkehub_FeatureMembershipConfigmanagementPolicyController struct {
	// Logs all denies and dry run failures.
	LogDeniesEnabled bool `json:"logDeniesEnabled,omitempty" yaml:"logDeniesEnabled,omitempty"`

	// Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: ["cloudmonitoring", "prometheus"]. Default: ["cloudmonitoring", "prometheus"]
	Monitoring Gkehub_FeatureMembershipConfigmanagementPolicyControllerMonitoring `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	// Enables mutation in policy controller. If true, mutation CRDs, webhook, and controller deployment will be deployed to the cluster.
	MutationEnabled bool `json:"mutationEnabled,omitempty" yaml:"mutationEnabled,omitempty"`

	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	ReferentialRulesEnabled bool `json:"referentialRulesEnabled,omitempty" yaml:"referentialRulesEnabled,omitempty"`

	// Installs the default template library along with Policy Controller.
	TemplateLibraryInstalled bool `json:"templateLibraryInstalled,omitempty" yaml:"templateLibraryInstalled,omitempty"`

	// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
	AuditIntervalSeconds string `json:"auditIntervalSeconds,omitempty" yaml:"auditIntervalSeconds,omitempty"`

	// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty" yaml:"exemptableNamespaces,omitempty"`
}
