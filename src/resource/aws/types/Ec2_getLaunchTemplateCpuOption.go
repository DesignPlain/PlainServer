package types

type Ec2_getLaunchTemplateCpuOption struct {
	//
	AmdSevSnp string `json:"amdSevSnp,omitempty" yaml:"amdSevSnp,omitempty"`

	//
	CoreCount int `json:"coreCount,omitempty" yaml:"coreCount,omitempty"`

	//
	ThreadsPerCore int `json:"threadsPerCore,omitempty" yaml:"threadsPerCore,omitempty"`
}
