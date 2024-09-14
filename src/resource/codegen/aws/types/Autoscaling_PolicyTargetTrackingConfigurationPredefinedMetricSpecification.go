package types

type Autoscaling_PolicyTargetTrackingConfigurationPredefinedMetricSpecification struct {
	// Metric type.
	PredefinedMetricType string `json:"predefinedMetricType,omitempty" yaml:"predefinedMetricType,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel string `json:"resourceLabel,omitempty" yaml:"resourceLabel,omitempty"`
}
