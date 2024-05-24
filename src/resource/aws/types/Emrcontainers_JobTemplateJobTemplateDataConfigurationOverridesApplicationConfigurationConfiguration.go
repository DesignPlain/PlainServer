package types

type Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesApplicationConfigurationConfiguration struct {
	// The classification within a configuration.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// A set of properties specified within a configuration classification.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
