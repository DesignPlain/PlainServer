package types

type Container_getClusterLoggingConfig struct {
	// GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
	EnableComponents []string `json:"enableComponents,omitempty" yaml:"enableComponents,omitempty"`
}
