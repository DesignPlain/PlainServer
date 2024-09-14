package types

type Dataplex_TaskSparkInfrastructureSpec struct {
	/*
	   Compute resources needed for a Task when using Dataproc Serverless.
	   Structure is documented below.
	*/
	Batch Dataplex_TaskSparkInfrastructureSpecBatch `json:"batch,omitempty" yaml:"batch,omitempty"`

	/*
	   Container Image Runtime Configuration.
	   Structure is documented below.
	*/
	ContainerImage Dataplex_TaskSparkInfrastructureSpecContainerImage `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`

	/*
	   Vpc network.
	   Structure is documented below.
	*/
	VpcNetwork Dataplex_TaskSparkInfrastructureSpecVpcNetwork `json:"vpcNetwork,omitempty" yaml:"vpcNetwork,omitempty"`
}
