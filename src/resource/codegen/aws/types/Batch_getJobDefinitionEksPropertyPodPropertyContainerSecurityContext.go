package types

type Batch_getJobDefinitionEksPropertyPodPropertyContainerSecurityContext struct {
	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the root user).
	Privileged bool `json:"privileged,omitempty" yaml:"privileged,omitempty"`

	//
	ReadOnlyRootFileSystem bool `json:"readOnlyRootFileSystem,omitempty" yaml:"readOnlyRootFileSystem,omitempty"`

	// When this parameter is specified, the container is run as the specified group ID (gid). If this parameter isn't specified, the default is the group that's specified in the image metadata.
	RunAsGroup int `json:"runAsGroup,omitempty" yaml:"runAsGroup,omitempty"`

	// When this parameter is specified, the container is run as a user with a uid other than 0. If this parameter isn't specified, so such rule is enforced.
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty" yaml:"runAsNonRoot,omitempty"`

	// When this parameter is specified, the container is run as the specified user ID (uid). If this parameter isn't specified, the default is the user that's specified in the image metadata.
	RunAsUser int `json:"runAsUser,omitempty" yaml:"runAsUser,omitempty"`
}
