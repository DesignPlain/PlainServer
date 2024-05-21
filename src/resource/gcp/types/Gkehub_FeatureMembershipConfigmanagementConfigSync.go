package types



type Gkehub_FeatureMembershipConfigmanagementConfigSync struct {
	/*
	   (Optional) Supported from ACM versions 1.12.0 onwards. Structure is documented below.

	   Use either `git` or `oci` config option.
	*/
	Oci Gkehub_FeatureMembershipConfigmanagementConfigSyncOci `json:"oci,omitempty" yaml:"oci,omitempty"`

	// Supported from ACM versions 1.10.0 onwards. Set to true to enable the Config Sync admission webhook to prevent drifts. If set to "false", disables the Config Sync admission webhook and does not prevent drifts.
	PreventDrift bool `json:"preventDrift,omitempty" yaml:"preventDrift,omitempty"`

	// Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode.
	SourceFormat string `json:"sourceFormat,omitempty" yaml:"sourceFormat,omitempty"`

	// (Optional) Structure is documented below.
	Git Gkehub_FeatureMembershipConfigmanagementConfigSyncGit `json:"git,omitempty" yaml:"git,omitempty"`

	// The Email of the Google Cloud Service Account (GSA) used for exporting Config Sync metrics to Cloud Monitoring. The GSA should have the Monitoring Metric Writer(roles/monitoring.metricWriter) IAM role. The Kubernetes ServiceAccount `default` in the namespace `config-management-monitoring` should be bound to the GSA.
	MetricsGcpServiceAccountEmail string `json:"metricsGcpServiceAccountEmail,omitempty" yaml:"metricsGcpServiceAccountEmail,omitempty"`
}
