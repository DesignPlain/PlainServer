package globalaccelerator

import types "DesignSphere_Server/src/resource/aws/types"

type EndpointGroup struct {
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol string `json:"healthCheckProtocol,omitempty" yaml:"healthCheckProtocol,omitempty"`

	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	PortOverrides []types.Globalaccelerator_EndpointGroupPortOverride `json:"portOverrides,omitempty" yaml:"portOverrides,omitempty"`

	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount int `json:"thresholdCount,omitempty" yaml:"thresholdCount,omitempty"`

	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage float64 `json:"trafficDialPercentage,omitempty" yaml:"trafficDialPercentage,omitempty"`

	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion string `json:"endpointGroupRegion,omitempty" yaml:"endpointGroupRegion,omitempty"`

	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPath string `json:"healthCheckPath,omitempty" yaml:"healthCheckPath,omitempty"`

	/*
	   The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	   the provider will only perform drift detection of its value when present in a configuration.
	*/
	HealthCheckPort int `json:"healthCheckPort,omitempty" yaml:"healthCheckPort,omitempty"`

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn string `json:"listenerArn,omitempty" yaml:"listenerArn,omitempty"`

	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations []types.Globalaccelerator_EndpointGroupEndpointConfiguration `json:"endpointConfigurations,omitempty" yaml:"endpointConfigurations,omitempty"`

	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds int `json:"healthCheckIntervalSeconds,omitempty" yaml:"healthCheckIntervalSeconds,omitempty"`
}
