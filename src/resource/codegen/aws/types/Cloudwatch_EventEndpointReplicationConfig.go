package types

type Cloudwatch_EventEndpointReplicationConfig struct {
	// The state of event replication. Valid values: `ENABLED`, `DISABLED`. The default state is `ENABLED`, which means you must supply a `role_arn`. If you don't have a `role_arn` or you don't want event replication enabled, set `state` to `DISABLED`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
