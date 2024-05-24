package types

type Evidently_LaunchMetricMonitorMetricDefinition struct {
	// Specifies the name for the metric.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies a label for the units that the metric is measuring.
	UnitLabel string `json:"unitLabel,omitempty" yaml:"unitLabel,omitempty"`

	// Specifies the value that is tracked to produce the metric.
	ValueKey string `json:"valueKey,omitempty" yaml:"valueKey,omitempty"`

	// Specifies the entity, such as a user or session, that does an action that causes a metric value to be recorded. An example is `userDetails.userID`.
	EntityIdKey string `json:"entityIdKey,omitempty" yaml:"entityIdKey,omitempty"`

	// Specifies The EventBridge event pattern that defines how the metric is recorded.
	EventPattern string `json:"eventPattern,omitempty" yaml:"eventPattern,omitempty"`
}
