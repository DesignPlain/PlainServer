package cloudwatch

import types "DesignSphere_Server/src/resource/aws/types"

type EventEndpoint struct {
	// A description of the global endpoint.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The event buses to use. The names of the event buses must be identical in each Region. Exactly two event buses are required. Documented below.
	EventBuses []types.Cloudwatch_EventEndpointEventBus `json:"eventBuses,omitempty" yaml:"eventBuses,omitempty"`

	// The name of the global endpoint.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Parameters used for replication. Documented below.
	ReplicationConfig types.Cloudwatch_EventEndpointReplicationConfig `json:"replicationConfig,omitempty" yaml:"replicationConfig,omitempty"`

	// The ARN of the IAM role used for replication between event buses.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Parameters used for routing, including the health check and secondary Region. Documented below.
	RoutingConfig types.Cloudwatch_EventEndpointRoutingConfig `json:"routingConfig,omitempty" yaml:"routingConfig,omitempty"`
}
