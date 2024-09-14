package types

type Batch_JobDefinitionEksPropertiesPodPropertiesContainers struct {
	// Entrypoint for the container. This isn't run within a shell. If this isn't specified, the ENTRYPOINT of the container image is used. Environment variable references are expanded using the container's environment.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// Docker image used to start the container.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Image pull policy for the container. Supported values are `Always`, `IfNotPresent`, and `Never`.
	ImagePullPolicy string `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`

	// Name of the container. If the name isn't specified, the default name "Default" is used. Each container in a pod must have a unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Security context for a job.
	SecurityContext Batch_JobDefinitionEksPropertiesPodPropertiesContainersSecurityContext `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`

	// Volume mounts for the container.
	VolumeMounts []Batch_JobDefinitionEksPropertiesPodPropertiesContainersVolumeMount `json:"volumeMounts,omitempty" yaml:"volumeMounts,omitempty"`

	// Array of arguments to the entrypoint. If this isn't specified, the CMD of the container image is used. This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes. Environment variable references are expanded using the container's environment.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// Type and amount of resources to assign to a container. The supported resources include `memory`, `cpu`, and `nvidia.com/gpu`.
	Resources Batch_JobDefinitionEksPropertiesPodPropertiesContainersResources `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Environment variables to pass to a container. See EKS Environment below.
	Envs []Batch_JobDefinitionEksPropertiesPodPropertiesContainersEnv `json:"envs,omitempty" yaml:"envs,omitempty"`
}
