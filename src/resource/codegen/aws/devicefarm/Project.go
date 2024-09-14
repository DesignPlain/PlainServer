package devicefarm

type Project struct {
	// Sets the execution timeout value (in minutes) for a project. All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
	DefaultJobTimeoutMinutes int `json:"defaultJobTimeoutMinutes,omitempty" yaml:"defaultJobTimeoutMinutes,omitempty"`

	// The name of the project
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
