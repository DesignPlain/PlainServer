package types

type Ec2transitgateway_getRouteTableRoutesFilter struct {
	// Set of values that are accepted for the given field.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayRoutes.html).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
