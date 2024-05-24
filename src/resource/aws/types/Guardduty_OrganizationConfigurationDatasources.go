package types

type Guardduty_OrganizationConfigurationDatasources struct {
	// Enable Kubernetes Audit Logs Monitoring automatically for new member accounts.
	Kubernetes Guardduty_OrganizationConfigurationDatasourcesKubernetes `json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`

	// Enable Malware Protection automatically for new member accounts.
	MalwareProtection Guardduty_OrganizationConfigurationDatasourcesMalwareProtection `json:"malwareProtection,omitempty" yaml:"malwareProtection,omitempty"`

	// Enable S3 Protection automatically for new member accounts.
	S3Logs Guardduty_OrganizationConfigurationDatasourcesS3Logs `json:"s3Logs,omitempty" yaml:"s3Logs,omitempty"`
}
