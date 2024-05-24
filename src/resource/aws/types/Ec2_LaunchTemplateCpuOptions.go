package types

type Ec2_LaunchTemplateCpuOptions struct {
	// The number of CPU cores for the instance.
	CoreCount int `json:"coreCount,omitempty" yaml:"coreCount,omitempty"`

	/*
	   The number of threads per CPU core.
	   To disable Intel Hyper-Threading Technology for the instance, specify a value of 1.
	   Otherwise, specify the default value of 2.

	   Both number of CPU cores and threads per core must be specified. Valid number of CPU cores and threads per core for the instance type can be found in the [CPU Options Documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html?shortFooter=true#cpu-options-supported-instances-values)
	*/
	ThreadsPerCore int `json:"threadsPerCore,omitempty" yaml:"threadsPerCore,omitempty"`

	// Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only. Valid values are `enabled` and `disabled`.
	AmdSevSnp string `json:"amdSevSnp,omitempty" yaml:"amdSevSnp,omitempty"`
}
