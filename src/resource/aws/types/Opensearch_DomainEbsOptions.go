package types

type Opensearch_DomainEbsOptions struct {
	// Type of EBS volumes attached to data nodes.
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Whether EBS volumes are attached to data nodes in the domain.
	EbsEnabled bool `json:"ebsEnabled,omitempty" yaml:"ebsEnabled,omitempty"`

	// Baseline input/output (I/O) performance of EBS volumes attached to data nodes. Applicable only for the GP3 and Provisioned IOPS EBS volume
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Specifies the throughput (in MiB/s) of the EBS volumes attached to data nodes. Applicable only for the gp3 volume type.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Size of EBS volumes attached to data nodes (in GiB).
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`
}
