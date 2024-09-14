package types

type Batch_getJobDefinitionEksPropertyPodPropertyContainer struct {
	// The image used to start a container.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// The image pull policy for the container.
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`

	// The name of the job definition to register. It can be up to 128 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The environment variables to pass to a container.  Array of EksContainerEnvironmentVariable objects.
	Envs []Batch_getJobDefinitionEksPropertyPodPropertyContainerEnv `json:"envs,omitempty" yaml:"envs,omitempty"`

	// The command that's passed to the container.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// The type and amount of resources to assign to a container.
	Resources []Batch_getJobDefinitionEksPropertyPodPropertyContainerResource `json:"resources,omitempty" yaml:"resources,omitempty"`

	// The security context for a job.
	SecurityContexts []Batch_getJobDefinitionEksPropertyPodPropertyContainerSecurityContext `json:"securityContexts,omitempty" yaml:"securityContexts,omitempty"`

	// The volume mounts for the container.
	VolumeMounts []Batch_getJobDefinitionEksPropertyPodPropertyContainerVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`

	// An array of arguments to the entrypoint
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`
}
