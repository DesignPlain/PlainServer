package types

type Dataplex_TaskNotebookInfrastructureSpec struct {
	/*
	   Compute resources needed for a Task when using Dataproc Serverless.
	   Structure is documented below.
	*/
	Batch Dataplex_TaskNotebookInfrastructureSpecBatch `json:"batch,omitempty" yaml:"batch,omitempty"`

	/*
	   Container Image Runtime Configuration.
	   Structure is documented below.
	*/
	ContainerImage Dataplex_TaskNotebookInfrastructureSpecContainerImage `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`

	/*
	   Vpc network.
	   Structure is documented below.
	*/
	VpcNetwork Dataplex_TaskNotebookInfrastructureSpecVpcNetwork `json:"vpcNetwork,omitempty" yaml:"vpcNetwork,omitempty"`
}
