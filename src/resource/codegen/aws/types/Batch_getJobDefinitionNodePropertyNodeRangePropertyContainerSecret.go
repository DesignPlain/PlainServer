package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerSecret struct {
	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The secret to expose to the container. The supported values are either the full Amazon Resource Name (ARN) of the AWS Secrets Manager secret or the full ARN of the parameter in the AWS Systems Manager Parameter Store.
	ValueFrom string `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}
