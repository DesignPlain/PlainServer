package costexplorer

type AnomalyMonitor struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The dimensions to evaluate. Valid values: `SERVICE`.
	MonitorDimension string `json:"monitorDimension,omitempty" yaml:"monitorDimension,omitempty"`

	// A valid JSON representation for the [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html) object.
	MonitorSpecification string `json:"monitorSpecification,omitempty" yaml:"monitorSpecification,omitempty"`

	// The possible type values. Valid values: `DIMENSIONAL` | `CUSTOM`.
	MonitorType string `json:"monitorType,omitempty" yaml:"monitorType,omitempty"`

	// The name of the monitor.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
