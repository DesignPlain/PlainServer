package types

type Container_ClusterLoggingConfig struct {
	/*
	   The GKE components exposing logs. Supported values include:
	   `SYSTEM_COMPONENTS`, `APISERVER`, `CONTROLLER_MANAGER`, `SCHEDULER`, and `WORKLOADS`.
	*/
	EnableComponents []string `json:"enableComponents,omitempty" yaml:"enableComponents,omitempty"`
}
