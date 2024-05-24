package types

type Emr_ClusterAutoTerminationPolicy struct {
	// Specifies the amount of idle time in seconds after which the cluster automatically terminates. You can specify a minimum of `60` seconds and a maximum of `604800` seconds (seven days).
	IdleTimeout int `json:"idleTimeout,omitempty" yaml:"idleTimeout,omitempty"`
}
