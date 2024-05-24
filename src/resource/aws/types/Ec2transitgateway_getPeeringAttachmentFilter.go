package types

type Ec2transitgateway_getPeeringAttachmentFilter struct {
	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayPeeringAttachments.html).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Set of values that are accepted for the given field.
	   An EC2 Transit Gateway Peering Attachment be selected if any one of the given values matches.
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
