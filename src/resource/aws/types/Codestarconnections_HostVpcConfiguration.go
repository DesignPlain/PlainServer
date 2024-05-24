package types

type Codestarconnections_HostVpcConfiguration struct {
	// ID of the security group or security groups associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ID of the subnet or subnets associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	TlsCertificate string `json:"tlsCertificate,omitempty" yaml:"tlsCertificate,omitempty"`

	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
