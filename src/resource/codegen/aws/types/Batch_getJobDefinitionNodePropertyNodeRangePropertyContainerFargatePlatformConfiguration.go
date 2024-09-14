package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerFargatePlatformConfiguration struct {
	// The AWS Fargate platform version where the jobs are running. A platform version is specified only for jobs that are running on Fargate resources.
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`
}
