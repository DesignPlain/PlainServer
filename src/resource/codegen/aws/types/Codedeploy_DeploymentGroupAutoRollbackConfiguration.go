package types

type Codedeploy_DeploymentGroupAutoRollbackConfiguration struct {
	// Indicates whether a defined automatic rollback configuration is currently enabled for this Deployment Group. If you enable automatic rollback, you must specify at least one event type.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   The event type or types that trigger a rollback. Supported types are `DEPLOYMENT_FAILURE`, `DEPLOYMENT_STOP_ON_ALARM` and `DEPLOYMENT_STOP_ON_REQUEST`.

	   _Only one `auto_rollback_configuration` is allowed_.
	*/
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`
}
