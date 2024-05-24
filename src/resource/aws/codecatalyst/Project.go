package codecatalyst

type Project struct {
	// The description of the project. This description will be displayed to all users of the project. We recommend providing a brief description of the project and its intended purpose.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The friendly name of the project that will be displayed to users.

	   The following arguments are optional:
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The name of the space.
	SpaceName string `json:"spaceName,omitempty" yaml:"spaceName,omitempty"`
}
