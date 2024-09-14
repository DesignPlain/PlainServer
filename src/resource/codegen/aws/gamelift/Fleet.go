package gamelift

import types "libds/aws/types"

type Fleet struct {
	// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
	FleetType string `json:"fleetType,omitempty" yaml:"fleetType,omitempty"`

	// ARN of an IAM role that instances in the fleet can assume.
	InstanceRoleArn string `json:"instanceRoleArn,omitempty" yaml:"instanceRoleArn,omitempty"`

	// ID of the GameLift Script to be deployed on the fleet.
	ScriptId string `json:"scriptId,omitempty" yaml:"scriptId,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Prompts GameLift to generate a TLS/SSL certificate for the fleet. See certificate_configuration.
	CertificateConfiguration types.Gamelift_FleetCertificateConfiguration `json:"certificateConfiguration,omitempty" yaml:"certificateConfiguration,omitempty"`

	// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
	Ec2InboundPermissions []types.Gamelift_FleetEc2InboundPermission `json:"ec2InboundPermissions,omitempty" yaml:"ec2InboundPermissions,omitempty"`

	// Name of an EC2 instance typeE.g., `t2.micro`
	Ec2InstanceType string `json:"ec2InstanceType,omitempty" yaml:"ec2InstanceType,omitempty"`

	// The name of the fleet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
	ResourceCreationLimitPolicy types.Gamelift_FleetResourceCreationLimitPolicy `json:"resourceCreationLimitPolicy,omitempty" yaml:"resourceCreationLimitPolicy,omitempty"`

	// ID of the GameLift Build to be deployed on the fleet.
	BuildId string `json:"buildId,omitempty" yaml:"buildId,omitempty"`

	// Game session protection policy to apply to all instances in this fleetE.g., `FullProtection`. Defaults to `NoProtection`.
	NewGameSessionProtectionPolicy string `json:"newGameSessionProtectionPolicy,omitempty" yaml:"newGameSessionProtectionPolicy,omitempty"`

	// Human-readable description of the fleet.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
	MetricGroups []string `json:"metricGroups,omitempty" yaml:"metricGroups,omitempty"`

	// Instructions for launching server processes on each instance in the fleet. See below.
	RuntimeConfiguration types.Gamelift_FleetRuntimeConfiguration `json:"runtimeConfiguration,omitempty" yaml:"runtimeConfiguration,omitempty"`
}
