package apprunner

import types "DesignSphere_Server/src/resource/aws/types"

type ObservabilityConfiguration struct {
	// Name of the observability configuration.
	ObservabilityConfigurationName string `json:"observabilityConfigurationName,omitempty" yaml:"observabilityConfigurationName,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing. See Trace Configuration below for more details.
	TraceConfiguration types.Apprunner_ObservabilityConfigurationTraceConfiguration `json:"traceConfiguration,omitempty" yaml:"traceConfiguration,omitempty"`
}
