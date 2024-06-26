package types

type Cloudrunv2_getServiceTemplateContainer struct {
	// List of environment variables to set in the container.
	Envs []Cloudrunv2_getServiceTemplateContainerEnv `json:"envs,omitempty" yaml:"envs,omitempty"`

	// URL of the Container image in Google Container Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Compute Resource requirements by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	Resources []Cloudrunv2_getServiceTemplateContainerResource `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.
	WorkingDir string `json:"workingDir,omitempty" yaml:"workingDir,omitempty"`

	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// Containers which should be started before this container. If specified the container will wait to start until all containers with the listed names are healthy.
	DependsOns []string `json:"dependsOns,omitempty" yaml:"dependsOns,omitempty"`

	// The name of the Cloud Run v2 Service.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible.

	   If omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on
	*/
	Ports []Cloudrunv2_getServiceTemplateContainerPort `json:"ports,omitempty" yaml:"ports,omitempty"`

	// Startup probe of application within the container. All other probes are disabled if a startup probe is provided, until it succeeds. Container will not be added to service endpoints if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	StartupProbes []Cloudrunv2_getServiceTemplateContainerStartupProbe `json:"startupProbes,omitempty" yaml:"startupProbes,omitempty"`

	// Volume to mount into the container's filesystem.
	VolumeMounts []Cloudrunv2_getServiceTemplateContainerVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`

	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// Periodic probe of container liveness. Container will be restarted if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	LivenessProbes []Cloudrunv2_getServiceTemplateContainerLivenessProbe `json:"livenessProbes,omitempty" yaml:"livenessProbes,omitempty"`
}
