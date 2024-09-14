package types

type Notebooks_EnvironmentContainerImage struct {
	/*
	   The path to the container image repository.
	   For example: gcr.io/{project_id}/{imageName}
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// The tag of the container image. If not specified, this defaults to the latest tag.
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`
}
