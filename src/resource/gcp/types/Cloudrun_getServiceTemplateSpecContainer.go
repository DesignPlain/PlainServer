package types

type Cloudrun_getServiceTemplateSpecContainer struct {
	/*
	   List of sources to populate environment variables in the container.
	   All invalid keys will be reported as an event when the container is starting.
	   When a key exists in multiple sources, the value associated with the last source will
	   take precedence. Values defined by an Env with a duplicate key will take
	   precedence.
	*/
	EnvFroms []Cloudrun_getServiceTemplateSpecContainerEnvFrom `json:"envFroms,omitempty" yaml:"envFroms,omitempty"`

	// List of environment variables to set in the container.
	Envs []Cloudrun_getServiceTemplateSpecContainerEnv `json:"envs,omitempty" yaml:"envs,omitempty"`

	// Compute Resources required by this container. Used to set values such as max memory
	Resources []Cloudrun_getServiceTemplateSpecContainerResource `json:"resources,omitempty" yaml:"resources,omitempty"`

	/*
	   Startup probe of application within the container.
	   All other probes are disabled if a startup probe is provided, until it
	   succeeds. Container will not be added to service endpoints if the probe fails.
	*/
	StartupProbes []Cloudrun_getServiceTemplateSpecContainerStartupProbe `json:"startupProbes,omitempty" yaml:"startupProbes,omitempty"`

	/*
	   Volume to mount into the container's filesystem.
	   Only supports SecretVolumeSources.
	*/
	VolumeMounts []Cloudrun_getServiceTemplateSpecContainerVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`

	/*
	   Container's working directory.
	   If not specified, the container runtime's default will be used, which
	   might be configured in the container image.
	*/
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`

	/*
	   Arguments to the entrypoint.
	   The docker image's CMD is used if this is not provided.
	*/
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	/*
	   Entrypoint array. Not executed within a shell.
	   The docker image's ENTRYPOINT is used if this is not provided.
	*/
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	/*
	   Docker image name. This is most often a reference to a container located
	   in the container registry, such as gcr.io/cloudrun/hello
	*/
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Periodic probe of container liveness. Container will be restarted if the probe fails.
	LivenessProbes []Cloudrun_getServiceTemplateSpecContainerLivenessProbe `json:"livenessProbes,omitempty" yaml:"livenessProbes,omitempty"`

	// The name of the Cloud Run Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of open ports in the container.
	Ports []Cloudrun_getServiceTemplateSpecContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`
}
