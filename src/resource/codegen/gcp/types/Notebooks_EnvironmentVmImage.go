package types

type Notebooks_EnvironmentVmImage struct {
	// Use this VM image family to find the image; the newest image in this family will be used.
	ImageFamily string `json:"imageFamily,omitempty" yaml:"imageFamily,omitempty"`

	// Use VM image name to find the image.
	ImageName string `json:"imageName,omitempty" yaml:"imageName,omitempty"`

	/*
	   The name of the Google Cloud project that this VM image belongs to.
	   Format: projects/{project_id}
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
