package types

type Batch_JobDefinitionEksPropertiesPodPropertiesContainers struct {
	// The security context for a job.
	SecurityContext Batch_JobDefinitionEksPropertiesPodPropertiesContainersSecurityContext `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`

	// The volume mounts for the container.
	VolumeMounts []Batch_JobDefinitionEksPropertiesPodPropertiesContainersVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`

	// The entrypoint for the container. This isn't run within a shell. If this isn't specified, the ENTRYPOINT of the container image is used. Environment variable references are expanded using the container's environment.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// The image pull policy for the container. Supported values are `Always`, `IfNotPresent`, and `Never`.
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`

	// The type and amount of resources to assign to a container. The supported resources include `memory`, `cpu`, and `nvidia.com/gpu`.
	Resources Batch_JobDefinitionEksPropertiesPodPropertiesContainersResources `json:"resources,omitempty" yaml:"resources,omitempty"`

	// The name of the container. If the name isn't specified, the default name "Default" is used. Each container in a pod must have a unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An array of arguments to the entrypoint. If this isn't specified, the CMD of the container image is used. This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes. Environment variable references are expanded using the container's environment.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// The environment variables to pass to a container. See EKS Environment below.
	Envs []Batch_JobDefinitionEksPropertiesPodPropertiesContainersEnv `json:"envs,omitempty" yaml:"envs,omitempty"`

	// The Docker image used to start the container.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
}
