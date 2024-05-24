package types

type Apigateway_getRestApiEndpointConfiguration struct {
	//
	Types []string `json:"types,omitempty" yaml:"types,omitempty"`

	//
	VpcEndpointIds []string `json:"vpcEndpointIds,omitempty" yaml:"vpcEndpointIds,omitempty"`
}
