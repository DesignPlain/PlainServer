package types

type Workbench_InstanceGceSetupVmImage struct {
	/*
	   Optional. Use this VM image family to find the image; the newest
	   image in this family will be used.
	*/
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	// Optional. Use VM image name to find the image.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name of the Google Cloud project that this VM image belongs to.
	   Format: {project_id}
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
