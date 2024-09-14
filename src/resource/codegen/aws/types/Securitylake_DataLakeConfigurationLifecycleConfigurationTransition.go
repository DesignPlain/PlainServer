package types

type Securitylake_DataLakeConfigurationLifecycleConfigurationTransition struct {
	// Number of days before data transition to a different S3 Storage Class in the Amazon Security Lake object.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}
