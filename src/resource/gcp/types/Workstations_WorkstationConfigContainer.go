package types

type Workstations_WorkstationConfigContainer struct {
	/*
	   Environment variables passed to the container.
	   The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
	*/
	Env map[string]string `json:"env,omitempty" yaml:"env,omitempty"`

	// Docker image defining the container. This image must be accessible by the config's service account.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// If set, overrides the USER specified in the image with the given uid.
	RunAsUser int `json:"runAsUser,omitempty" yaml:"runAsUser,omitempty"`

	// If set, overrides the default DIR specified by the image.
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`

	// Arguments passed to the entrypoint.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// If set, overrides the default ENTRYPOINT specified by the image.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`
}
