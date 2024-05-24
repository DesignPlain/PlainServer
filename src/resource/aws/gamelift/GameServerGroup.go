package gamelift

import types "DesignSphere_Server/src/resource/aws/types"

type GameServerGroup struct {
	// Key-value map of resource tags
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Indicates whether instances in the game server group are protected from early termination.
	   Unprotected instances that have active game servers running might be terminated during a scale-down event,
	   causing players to be dropped from the game.
	   Protected instances cannot be terminated while there are active game servers running except in the event
	   of a forced game server group deletion.
	   Valid values: `NO_PROTECTION`, `FULL_PROTECTION`. Defaults to `NO_PROTECTION`.
	*/
	GameServerProtectionPolicy string `json:"gameServerProtectionPolicy,omitempty" yaml:"gameServerProtectionPolicy,omitempty"`

	//
	InstanceDefinitions []types.Gamelift_GameServerGroupInstanceDefinition `json:"instanceDefinitions,omitempty" yaml:"instanceDefinitions,omitempty"`

	/*
	   The maximum number of instances allowed in the EC2 Auto Scaling group.
	   During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
	*/
	MaxSize int `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`

	/*
	   The minimum number of instances allowed in the EC2 Auto Scaling group.
	   During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
	*/
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	//
	AutoScalingPolicy types.Gamelift_GameServerGroupAutoScalingPolicy `json:"autoScalingPolicy,omitempty" yaml:"autoScalingPolicy,omitempty"`

	/*
	   Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
	   Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
	*/
	BalancingStrategy string `json:"balancingStrategy,omitempty" yaml:"balancingStrategy,omitempty"`

	/*
	   Name of the game server group.
	   This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
	*/
	GameServerGroupName string `json:"gameServerGroupName,omitempty" yaml:"gameServerGroupName,omitempty"`

	//
	LaunchTemplate types.Gamelift_GameServerGroupLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`

	/*
	   A list of VPC subnets to use with instances in the game server group.
	   By default, all GameLift FleetIQ-supported Availability Zones are used.
	*/
	VpcSubnets []string `json:"vpcSubnets,omitempty" yaml:"vpcSubnets,omitempty"`
}
