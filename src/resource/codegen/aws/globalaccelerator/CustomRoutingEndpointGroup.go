package globalaccelerator

import types "libds/aws/types"

type CustomRoutingEndpointGroup struct {
	// The port ranges and protocols for all endpoints in a custom routing endpoint group to accept client traffic on. Fields documented below.
	DestinationConfigurations []types.Globalaccelerator_CustomRoutingEndpointGroupDestinationConfiguration `json:"destinationConfigurations,omitempty" yaml:"destinationConfigurations,omitempty"`

	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations []types.Globalaccelerator_CustomRoutingEndpointGroupEndpointConfiguration `json:"endpointConfigurations,omitempty" yaml:"endpointConfigurations,omitempty"`

	// The name of the AWS Region where the custom routing endpoint group is located.
	EndpointGroupRegion string `json:"endpointGroupRegion,omitempty" yaml:"endpointGroupRegion,omitempty"`

	// The Amazon Resource Name (ARN) of the custom routing listener.
	ListenerArn string `json:"listenerArn,omitempty" yaml:"listenerArn,omitempty"`
}
