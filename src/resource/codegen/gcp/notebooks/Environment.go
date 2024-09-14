package notebooks

import types "libds/gcp/types"

type Environment struct {
	// A brief description of this environment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Display name of this environment for the UI.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   A reference to the zone where the machine resides.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The name specified for the Environment instance.
	   Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Path to a Bash script that automatically runs after a notebook instance fully boots up.
	   The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	*/
	PostStartupScript string `json:"postStartupScript,omitempty" yaml:"postStartupScript,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Use a Compute Engine VM image to start the notebook instance.
	   Structure is documented below.
	*/
	VmImage types.Notebooks_EnvironmentVmImage `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`

	/*
	   Use a container image to start the notebook instance.
	   Structure is documented below.
	*/
	ContainerImage types.Notebooks_EnvironmentContainerImage `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`
}
