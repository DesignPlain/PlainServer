package codecatalyst

type SourceRepository struct {
	// The description of the project. This description will be displayed to all users of the project. We recommend providing a brief description of the project and its intended purpose.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the source repository. For more information about name requirements, see [Quotas for source repositories](https://docs.aws.amazon.com/codecatalyst/latest/userguide/source-quotas.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name of the project in the CodeCatalyst space.

	   The following arguments are optional:
	*/
	ProjectName string `json:"projectName,omitempty" yaml:"projectName,omitempty"`

	// The name of the CodeCatalyst space.
	SpaceName string `json:"spaceName,omitempty" yaml:"spaceName,omitempty"`
}
