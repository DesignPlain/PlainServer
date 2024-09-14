package types

type Container_getClusterResourceUsageExportConfig struct {
	// Whether to enable resource consumption metering on this cluster. When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true.
	EnableResourceConsumptionMetering bool `json:"enableResourceConsumptionMetering,omitempty" yaml:"enableResourceConsumptionMetering,omitempty"`

	// Parameters for using BigQuery as the destination of resource usage export.
	BigqueryDestinations []Container_getClusterResourceUsageExportConfigBigqueryDestination `json:"bigqueryDestinations,omitempty" yaml:"bigqueryDestinations,omitempty"`

	// Whether to enable network egress metering for this cluster. If enabled, a daemonset will be created in the cluster to meter network egress traffic.
	EnableNetworkEgressMetering bool `json:"enableNetworkEgressMetering,omitempty" yaml:"enableNetworkEgressMetering,omitempty"`
}
