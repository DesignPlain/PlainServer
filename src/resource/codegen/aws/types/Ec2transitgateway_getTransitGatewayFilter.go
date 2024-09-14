package types

type Ec2transitgateway_getTransitGatewayFilter struct {
	// Name of the field to filter by, as defined by the [underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGateways.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of one or more values for the filter.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
