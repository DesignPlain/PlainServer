package types

type Cloudrun_ServiceTemplateSpecContainer struct {
	/*
	   Entrypoint array. Not executed within a shell.
	   The docker image's ENTRYPOINT is used if this is not provided.
	*/
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	/*
	   List of environment variables to set in the container.
	   Structure is documented below.
	*/
	Envs []Cloudrun_ServiceTemplateSpecContainerEnv `json:"envs,omitempty" yaml:"envs,omitempty"`

	/*
	   (Optional, Deprecated)
	   Container's working directory.
	   If not specified, the container runtime's default will be used, which
	   might be configured in the container image.

	   > --Warning:-- `working_dir` is deprecated and will be removed in a future major release. This field is not supported by the Cloud Run API.
	*/
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`

	/*
	   Startup probe of application within the container.
	   All other probes are disabled if a startup probe is provided, until it
	   succeeds. Container will not be added to service endpoints if the probe fails.
	   Structure is documented below.
	*/
	StartupProbe Cloudrun_ServiceTemplateSpecContainerStartupProbe `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`

	/*
	   Arguments to the entrypoint.
	   The docker image's CMD is used if this is not provided.
	*/
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	/*
	   (Optional, Deprecated)
	   List of sources to populate environment variables in the container.
	   All invalid keys will be reported as an event when the container is starting.
	   When a key exists in multiple sources, the value associated with the last source will
	   take precedence. Values defined by an Env with a duplicate key will take
	   precedence.
	   Structure is documented below.

	   > --Warning:-- `env_from` is deprecated and will be removed in a future major release. This field is not supported by the Cloud Run API.
	*/
	EnvFroms []Cloudrun_ServiceTemplateSpecContainerEnvFrom `json:"envFroms,omitempty" yaml:"envFroms,omitempty"`

	/*
	   Docker image name. This is most often a reference to a container located
	   in the container registry, such as gcr.io/cloudrun/hello
	*/
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	/*
	   Periodic probe of container liveness. Container will be restarted if the probe fails. More info:
	   https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	   Structure is documented below.
	*/
	LivenessProbe Cloudrun_ServiceTemplateSpecContainerLivenessProbe `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`

	// Name of the container
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   List of open ports in the container.
	   Structure is documented below.
	*/
	Ports []Cloudrun_ServiceTemplateSpecContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`

	/*
	   Compute Resources required by this container. Used to set values such as max memory
	   Structure is documented below.
	*/
	Resources Cloudrun_ServiceTemplateSpecContainerResources `json:"resources,omitempty" yaml:"resources,omitempty"`

	/*
	   Volume to mount into the container's filesystem.
	   Only supports SecretVolumeSources.
	   Structure is documented below.
	*/
	VolumeMounts []Cloudrun_ServiceTemplateSpecContainerVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`
}
