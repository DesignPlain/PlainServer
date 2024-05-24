package types

type S3_AccessPointVpcConfiguration struct {
	// This access point will only allow connections from the specified VPC ID.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
