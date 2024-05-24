package types

type Ec2transitgateway_getAttachmentsFilter struct {
	// Name of the filter check available value on [official documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of one or more values for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
