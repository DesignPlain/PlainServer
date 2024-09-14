package types

type Gamelift_GameServerGroupAutoScalingPolicy struct {
	/*
	   Length of time, in seconds, it takes for a new instance to start
	   new game server processes and register with GameLift FleetIQ.
	   Specifying a warm-up time can be useful, particularly with game servers that take a long time to start up,
	   because it avoids prematurely starting new instances. Defaults to `60`.
	*/
	EstimatedInstanceWarmup int `json:"estimatedInstanceWarmup,omitempty" yaml:"estimatedInstanceWarmup,omitempty"`

	//
	TargetTrackingConfiguration Gamelift_GameServerGroupAutoScalingPolicyTargetTrackingConfiguration `json:"targetTrackingConfiguration,omitempty" yaml:"targetTrackingConfiguration,omitempty"`
}
