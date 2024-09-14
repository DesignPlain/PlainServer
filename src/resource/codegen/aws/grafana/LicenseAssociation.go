package grafana

type LicenseAssociation struct {
	// The workspace id.
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`

	// A token from Grafana Labs that ties your AWS account with a Grafana Labs account.
	GrafanaToken string `json:"grafanaToken,omitempty" yaml:"grafanaToken,omitempty"`

	// The type of license for the workspace license association. Valid values are `ENTERPRISE` and `ENTERPRISE_FREE_TRIAL`.
	LicenseType string `json:"licenseType,omitempty" yaml:"licenseType,omitempty"`
}
