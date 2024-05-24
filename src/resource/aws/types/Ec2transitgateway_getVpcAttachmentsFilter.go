package types

type Ec2transitgateway_getVpcAttachmentsFilter struct {
	// List of one or more values for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Name of the filter check available value on [official documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayVpcAttachments.html)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
