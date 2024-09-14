package monitoring

type MonitoredProject struct {
	/*
	   Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}


	   - - -
	*/
	MetricsScope string `json:"metricsScope,omitempty" yaml:"metricsScope,omitempty"`

	// Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
