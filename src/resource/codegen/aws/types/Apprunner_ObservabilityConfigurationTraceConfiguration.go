package types

type Apprunner_ObservabilityConfigurationTraceConfiguration struct {
	// Implementation provider chosen for tracing App Runner services. Valid values: `AWSXRAY`.
	Vendor string `json:"vendor,omitempty" yaml:"vendor,omitempty"`
}
