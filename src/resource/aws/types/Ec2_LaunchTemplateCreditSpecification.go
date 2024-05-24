package types

type Ec2_LaunchTemplateCreditSpecification struct {
	/*
	   The credit option for CPU usage.
	   Can be `standard` or `unlimited`.
	   T3 instances are launched as `unlimited` by default.
	   T2 instances are launched as `standard` by default.
	*/
	CpuCredits string `json:"cpuCredits,omitempty" yaml:"cpuCredits,omitempty"`
}
