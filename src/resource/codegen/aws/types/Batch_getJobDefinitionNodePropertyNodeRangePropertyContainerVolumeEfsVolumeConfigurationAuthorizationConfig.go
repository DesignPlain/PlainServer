package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolumeEfsVolumeConfigurationAuthorizationConfig struct {
	// The Amazon EFS access point ID to use.
	AccessPointId string `json:"accessPointId,omitempty" yaml:"accessPointId,omitempty"`

	// Whether or not to use the AWS Batch job IAM role defined in a job definition when mounting the Amazon EFS file system.
	Iam string `json:"iam,omitempty" yaml:"iam,omitempty"`
}
