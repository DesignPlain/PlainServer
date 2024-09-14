package types

type Quicksight_DataSetRowLevelPermissionDataSet struct {
	// Type of permissions to use when interpreting the permissions for RLS. Valid values are `GRANT_ACCESS` and `DENY_ACCESS`.
	PermissionPolicy string `json:"permissionPolicy,omitempty" yaml:"permissionPolicy,omitempty"`

	// Status of the row-level security permission dataset. If enabled, the status is `ENABLED`. If disabled, the status is `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// ARN of the dataset that contains permissions for RLS.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// User or group rules associated with the dataset that contains permissions for RLS.
	FormatVersion string `json:"formatVersion,omitempty" yaml:"formatVersion,omitempty"`

	// Namespace associated with the dataset that contains permissions for RLS.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
