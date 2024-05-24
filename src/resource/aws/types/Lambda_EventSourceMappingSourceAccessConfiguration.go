package types

type Lambda_EventSourceMappingSourceAccessConfiguration struct {
	// The type of authentication protocol, VPC components, or virtual host for your event source. For valid values, refer to the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/api/API_SourceAccessConfiguration.html).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The URI for this configuration.  For type `VPC_SUBNET` the value should be `subnet:subnet_id` where `subnet_id` is the value you would find in an aws.ec2.Subnet resource's id attribute.  For type `VPC_SECURITY_GROUP` the value should be `security_group:security_group_id` where `security_group_id` is the value you would find in an aws.ec2.SecurityGroup resource's id attribute.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
