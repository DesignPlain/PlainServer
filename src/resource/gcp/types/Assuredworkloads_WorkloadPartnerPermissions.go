package types

type Assuredworkloads_WorkloadPartnerPermissions struct {
	// Allow the partner to view inspectability logs and monitoring violations.
	DataLogsViewer bool `json:"dataLogsViewer,omitempty" yaml:"dataLogsViewer,omitempty"`

	// Optional. Allow partner to view access approval logs.
	ServiceAccessApprover bool `json:"serviceAccessApprover,omitempty" yaml:"serviceAccessApprover,omitempty"`

	// Optional. Allow partner to view violation alerts.
	AssuredWorkloadsMonitoring bool `json:"assuredWorkloadsMonitoring,omitempty" yaml:"assuredWorkloadsMonitoring,omitempty"`
}
