package cloud9

type EnvironmentEC2 struct {
	// The number of minutes until the running instance is shut down after the environment has last been used.
	AutomaticStopTimeMinutes int `json:"automaticStopTimeMinutes,omitempty" yaml:"automaticStopTimeMinutes,omitempty"`

	/*
	   The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
	   - `amazonlinux-2-x86_64`
	   - `amazonlinux-2023-x86_64`
	   - `ubuntu-18.04-x86_64`
	   - `ubuntu-22.04-x86_64`
	   - `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
	   - `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2023-x86_64`
	   - `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
	   - `resolve:ssm:/aws/service/cloud9/amis/ubuntu-22.04-x86_64`
	*/
	ImageId string `json:"imageId,omitempty" yaml:"imageId,omitempty"`

	// The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
	OwnerArn string `json:"ownerArn,omitempty" yaml:"ownerArn,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
	ConnectionType string `json:"connectionType,omitempty" yaml:"connectionType,omitempty"`

	// The description of the environment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The type of instance to connect to the environment, e.g., `t2.micro`.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The name of the environment.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
