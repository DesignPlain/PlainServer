package types

type Dataplex_TaskSparkInfrastructureSpecVpcNetwork struct {
	// List of network tags to apply to the job.
	NetworkTags []string `json:"networkTags,omitempty" yaml:"networkTags,omitempty"`

	// The Cloud VPC sub-network in which the job is run.
	SubNetwork string `json:"subNetwork,omitempty" yaml:"subNetwork,omitempty"`

	// The Cloud VPC network in which the job is run. By default, the Cloud VPC network named Default within the project is used.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
