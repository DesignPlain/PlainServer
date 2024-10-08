package types

type Cloudrunv2_JobTemplateTemplateContainer struct {
	/*
	   List of environment variables to set in the container.
	   Structure is documented below.
	*/
	Envs []Cloudrunv2_JobTemplateTemplateContainerEnv `json:"envs,omitempty" yaml:"envs,omitempty"`

	// URL of the Container image in Google Container Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	/*
	   Compute Resource requirements by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	   Structure is documented below.
	*/
	Resources Cloudrunv2_JobTemplateTemplateContainerResources `json:"resources,omitempty" yaml:"resources,omitempty"`

	/*
	   Volume to mount into the container's filesystem.
	   Structure is documented below.
	*/
	VolumeMounts []Cloudrunv2_JobTemplateTemplateContainerVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`

	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// Name of the container specified as a DNS_LABEL.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible.
	   If omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on
	   Structure is documented below.
	*/
	Ports []Cloudrunv2_JobTemplateTemplateContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`

	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`
}
