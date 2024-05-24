package types

type Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig struct {
	// A list of security group IDs to associate with Kinesis Firehose.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of subnet IDs to associate with Kinesis Firehose.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	//
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The ARN of the IAM role to be assumed by Firehose for calling the Amazon EC2 configuration API and for creating network interfaces. Make sure role has necessary [IAM permissions](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-es-vpc)
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
