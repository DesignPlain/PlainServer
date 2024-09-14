package types

type Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesApplicationConfiguration struct {
	// A set of properties specified within a configuration classification.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// The classification within a configuration.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// A list of additional configurations to apply within a configuration object.
	Configurations []Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesApplicationConfigurationConfiguration `json:"configurations,omitempty" yaml:"configurations,omitempty"`
}
