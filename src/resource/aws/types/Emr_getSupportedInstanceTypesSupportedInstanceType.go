package types

type Emr_getSupportedInstanceTypesSupportedInstanceType struct {
	// CPU architecture.
	Architecture string `json:"architecture,omitempty" yaml:"architecture,omitempty"`

	// Indicates whether the instance type supports Amazon EBS optimization.
	EbsOptimizedAvailable bool `json:"ebsOptimizedAvailable,omitempty" yaml:"ebsOptimizedAvailable,omitempty"`

	// Memory that is available to Amazon EMR from the instance type.
	MemoryGb float64 `json:"memoryGb,omitempty" yaml:"memoryGb,omitempty"`

	// Number of disks for the instance type.
	NumberOfDisks int `json:"numberOfDisks,omitempty" yaml:"numberOfDisks,omitempty"`

	// Storage capacity of the instance type.
	StorageGb int `json:"storageGb,omitempty" yaml:"storageGb,omitempty"`

	// Amazon EC2 instance type. For example, `m5.xlarge`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Indicates whether the instance type uses Amazon EBS optimization by default.
	EbsOptimizedByDefault bool `json:"ebsOptimizedByDefault,omitempty" yaml:"ebsOptimizedByDefault,omitempty"`

	// Indicates whether the instance type only supports Amazon EBS.
	EbsStorageOnly bool `json:"ebsStorageOnly,omitempty" yaml:"ebsStorageOnly,omitempty"`

	// The Amazon EC2 family and generation for the instance type.
	InstanceFamilyId string `json:"instanceFamilyId,omitempty" yaml:"instanceFamilyId,omitempty"`

	// Indicates whether the instance type only supports 64-bit architecture.
	Is64BitsOnly bool `json:"is64BitsOnly,omitempty" yaml:"is64BitsOnly,omitempty"`

	// The number of vCPUs available for the instance type.
	Vcpu int `json:"vcpu,omitempty" yaml:"vcpu,omitempty"`
}
