package types

type Projects_getProjectProject struct {
	// Creation time in RFC3339 UTC "Zulu" format.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// A set of key/value label pairs assigned on a project.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The Project lifecycle state.
	LifecycleState string `json:"lifecycleState,omitempty" yaml:"lifecycleState,omitempty"`

	// The optional user-assigned display name of the project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The numeric identifier of the project.
	Number string `json:"number,omitempty" yaml:"number,omitempty"`

	// An optional reference to a parent resource.
	Parent map[string]string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// The project id of the project.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
