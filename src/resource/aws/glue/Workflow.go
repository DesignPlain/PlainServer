package glue

type Workflow struct {
	// The name you assign to this workflow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A map of default run properties for this workflow. These properties are passed to all jobs associated to the workflow.
	DefaultRunProperties map[string]string `json:"defaultRunProperties,omitempty" yaml:"defaultRunProperties,omitempty"`

	// Description of the workflow.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Prevents exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
	MaxConcurrentRuns int `json:"maxConcurrentRuns,omitempty" yaml:"maxConcurrentRuns,omitempty"`
}
