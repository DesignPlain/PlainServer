package types

type Autoscaling_PolicyTargetTrackingConfiguration struct {
	// Target value for the metric.
	TargetValue float64 `json:"targetValue,omitempty" yaml:"targetValue,omitempty"`

	// Customized metric. Conflicts with `predefined_metric_specification`.
	CustomizedMetricSpecification Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty" yaml:"customizedMetricSpecification,omitempty"`

	// Whether scale in by the target tracking policy is disabled.
	DisableScaleIn bool `json:"disableScaleIn,omitempty" yaml:"disableScaleIn,omitempty"`

	// Predefined metric. Conflicts with `customized_metric_specification`.
	PredefinedMetricSpecification Autoscaling_PolicyTargetTrackingConfigurationPredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty" yaml:"predefinedMetricSpecification,omitempty"`
}
