package types

type Ec2_InstanceCreditSpecification struct {
	// Credit option for CPU usage. Valid values include `standard` or `unlimited`. T3 instances are launched as unlimited by default. T2 instances are launched as standard by default.
	CpuCredits string `json:"cpuCredits,omitempty" yaml:"cpuCredits,omitempty"`
}
