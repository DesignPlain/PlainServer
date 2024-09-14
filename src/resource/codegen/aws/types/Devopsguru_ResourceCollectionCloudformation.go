package types

type Devopsguru_ResourceCollectionCloudformation struct {
	// Array of the names of the AWS CloudFormation stacks. If `type` is `AWS_SERVICE` (all acccount resources) this array should be a single item containing a wildcard (`"-"`).
	StackNames []string `json:"stackNames,omitempty" yaml:"stackNames,omitempty"`
}
