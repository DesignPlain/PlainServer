package types

type Appsync_GraphQLApiEnhancedMetricsConfig struct {
	// How data source metrics will be emitted to CloudWatch. Valid values: `FULL_REQUEST_DATA_SOURCE_METRICS`, `PER_DATA_SOURCE_METRICS`
	DataSourceLevelMetricsBehavior string `json:"dataSourceLevelMetricsBehavior,omitempty" yaml:"dataSourceLevelMetricsBehavior,omitempty"`

	// How operation metrics will be emitted to CloudWatch. Valid values: `ENABLED`, `DISABLED`
	OperationLevelMetricsConfig string `json:"operationLevelMetricsConfig,omitempty" yaml:"operationLevelMetricsConfig,omitempty"`

	// How resolver metrics will be emitted to CloudWatch. Valid values: `FULL_REQUEST_RESOLVER_METRICS`, `PER_RESOLVER_METRICS`
	ResolverLevelMetricsBehavior string `json:"resolverLevelMetricsBehavior,omitempty" yaml:"resolverLevelMetricsBehavior,omitempty"`
}
