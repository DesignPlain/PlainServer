package grafana

type LicenseAssociation struct {
	// The type of license for the workspace license association. Valid values are `ENTERPRISE` and `ENTERPRISE_FREE_TRIAL`.
	LicenseType string `json:"licenseType,omitempty" yaml:"licenseType,omitempty"`

	// The workspace id.
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
}
