package types

type Devopsguru_getResourceCollectionCloudformation struct {
	// Array of the names of the AWS CloudFormation stacks.
	StackNames []string `json:"stackNames,omitempty" yaml:"stackNames,omitempty"`
}
